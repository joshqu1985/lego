package pkg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/iancoleman/strcase"

	"github.com/joshqu1985/lego/cmd/internal/pkg/g4"
)

const (
	ProtocolComment = 1
)

type ProtoVisitor struct {
	g4.Protobuf3Visitor
	Tree Tree
}

var BasicTypes = []string{
	"bool",
	"string",
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"float32",
	"float64",
	"byte",
	"rune",
	"complex64",
	"complex128",
}

func Parse(filename string) (*Tree, error) {
	if filename == "" {
		return nil, errors.New("filename is empty")
	}

	file, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, err
	}

	stream := antlr.NewCommonTokenStream(g4.NewProtobuf3Lexer(file),
		antlr.TokenDefaultChannel)

	data := g4.NewProtobuf3Parser(stream).Proto().Accept(newProtoVisitor())

	tree, _ := data.(*Tree)
	for _, node := range tree.Structs {
		preprocessStruct(node)
	}

	return tree, nil
}

func preprocessStruct(item *StructNode) {
	fieldTag, ok := item.Options["message_tag"]
	if !ok {
		fieldTag = "json"
	}
	tags := strings.Split(fieldTag, ",")

	for _, field := range item.Fields {
		field.Name = strcase.ToCamel(field.Name)
		field.Type = preprocessType(field.Type, field.Repeated)

		vals := make([]string, 0)
		for _, tag := range tags {
			vals = append(vals, fmt.Sprintf("%s:%q", tag, strcase.ToSnake(field.Name))) // nolint:goconst
		}
		field.Tag = "`" + strings.Join(vals, " ") + "`"
	}

	for _, field := range item.MapFields {
		field.Name = strcase.ToCamel(field.Name)
		vals := make([]string, 0)
		for _, tag := range tags {
			vals = append(vals, fmt.Sprintf("%s:%q", tag, strcase.ToSnake(field.Name)))
		}
		field.Tag = "`" + strings.Join(vals, " ") + "`"
	}
}

func preprocessType(t string, repeated bool) string {
	if isBasicType(t) {
		if repeated {
			return "[]" + t
		}

		return t
	}

	if repeated {
		return "[]*" + t
	} else {
		return "*" + t
	}
}

func newProtoVisitor() *ProtoVisitor {
	return &ProtoVisitor{
		Tree: Tree{Options: make(map[string]string)},
	}
}

func (v *ProtoVisitor) VisitProto(ctx *g4.ProtoContext) any {
	for _, each := range ctx.AllImportStatement() {
		if statement := each.Accept(v); statement != nil {
			iStatement, _ := statement.(string)
			v.Tree.Imports = append(v.Tree.Imports, iStatement)
		}
	}

	for _, each := range ctx.AllOptionStatement() {
		option, _ := each.Accept(v).(*OptionConst)
		if option != nil {
			v.Tree.Options[option.Name] = option.Constant
		}
	}

	for _, each := range ctx.AllTopLevelDef() {
		each.Accept(v)
	}

	if packageStatement := ctx.PackageStatement(0); packageStatement != nil {
		v.Tree.Package = packageStatement.FullIdent().GetText()
	}

	return &v.Tree
}

func (v *ProtoVisitor) VisitImportStatement(ctx *g4.ImportStatementContext) any {
	if ctx.StrLit() == nil {
		return nil
	}

	return ctx.StrLit().Accept(v)
}

func (v *ProtoVisitor) VisitOptionStatement(ctx *g4.OptionStatementContext) any {
	if ctx.OptionName() == nil || ctx.Constant() == nil {
		return nil
	}

	name, _ := ctx.OptionName().Accept(v).(string)
	name = strings.TrimFunc(name, func(r rune) bool { return r == '(' || r == ')' })

	constant, _ := ctx.Constant().Accept(v).(string)
	constant = strings.TrimFunc(constant, func(r rune) bool { return r == '"' })

	return &OptionConst{
		Name:     name,
		Constant: constant,
	}
}

func (v *ProtoVisitor) VisitTopLevelDef(ctx *g4.TopLevelDefContext) any {
	if ctx.EnumDef() != nil {
		if enum := ctx.EnumDef().Accept(v); enum != nil {
			ienum, _ := enum.(*EnumNode)
			v.Tree.Enums = append(v.Tree.Enums, ienum)
		}
	}

	if ctx.MessageDef() != nil {
		if message := ctx.MessageDef().Accept(v); message != nil {
			imessage, _ := message.(*StructNode)
			v.Tree.Structs = append(v.Tree.Structs, imessage)
		}
	}

	if ctx.ServiceDef() != nil {
		if service := ctx.ServiceDef().Accept(v); service != nil {
			iservice, _ := service.(*ServiceNode)
			v.Tree.Services = append(v.Tree.Services, iservice)
		}
	}

	return v
}

func (v *ProtoVisitor) VisitMessageDef(ctx *g4.MessageDefContext) any {
	message := &StructNode{
		Document: v.getHiddenTokensToLeft(ctx, ProtocolComment),
		Options:  make(map[string]string),
	}

	if ctx.MessageName() != nil {
		if name := ctx.MessageName().Accept(v); name != nil {
			message.Name, _ = name.(string)
		}
	}

	if ctx.MessageBody() != nil {
		body, _ := ctx.MessageBody().Accept(v).(*StructNode)
		if body != nil {
			for key, val := range body.Options {
				message.Options[key] = val
			}
			message.Fields = append(message.Fields, body.Fields...)
			message.MapFields = append(message.MapFields, body.MapFields...)
		}
	}

	return message
}

func (v *ProtoVisitor) VisitMessageName(ctx *g4.MessageNameContext) any {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitMessageBody(ctx *g4.MessageBodyContext) any {
	elements := &StructNode{
		Options: make(map[string]string),
	}

	for _, each := range ctx.AllMessageElement() {
		element := each.Accept(v)
		if element == nil {
			continue
		}

		if field, ok := element.(*StructField); ok {
			elements.Fields = append(elements.Fields, field)

			continue
		}

		if field, ok := element.(*StructMap); ok {
			elements.MapFields = append(elements.MapFields, field)

			continue
		}

		if field, ok := element.(*OptionConst); ok {
			elements.Options[field.Name] = field.Constant
		}
	}

	return elements
}

func (v *ProtoVisitor) VisitMessageElement(ctx *g4.MessageElementContext) any {
	if ctx.MessageDef() != nil {
		if message := ctx.MessageDef().Accept(v); message != nil {
			imessage, _ := message.(*StructNode)
			v.Tree.Structs = append(v.Tree.Structs, imessage)
		}
	}

	if ctx.EnumDef() != nil {
		if enum := ctx.EnumDef().Accept(v); enum != nil {
			ienum, _ := enum.(*EnumNode)
			v.Tree.Enums = append(v.Tree.Enums, ienum)
		}
	}

	if ctx.Field() != nil {
		if field := ctx.Field().Accept(v); field != nil {
			return field
		}
	}

	if ctx.MapField() != nil {
		if field := ctx.MapField().Accept(v); field != nil {
			return field
		}
	}

	if ctx.OptionStatement() != nil {
		if statement := ctx.OptionStatement().Accept(v); statement != nil {
			return statement
		}
	}

	return nil
}

func (v *ProtoVisitor) VisitField(ctx *g4.FieldContext) any {
	field := &StructField{
		Comment: v.getHiddenTokensToRight(ctx, ProtocolComment),
	}

	if ctx.FieldLabel() != nil && ctx.FieldLabel().REPEATED() != nil {
		if ctx.FieldLabel().REPEATED().GetSymbol().GetText() == "repeated" {
			field.Repeated = true
		}
	}

	if ctx.Type_() != nil {
		if xtype := ctx.Type_().Accept(v); xtype != nil {
			field.Type, _ = xtype.(string)
		}
	}

	if ctx.FieldName() != nil {
		if name := ctx.FieldName().Accept(v); name != nil {
			field.Name, _ = name.(string)
		}
	}

	return field
}

func (v *ProtoVisitor) VisitMapField(ctx *g4.MapFieldContext) any {
	field := &StructMap{
		Comment: v.getHiddenTokensToRight(ctx, ProtocolComment),
	}

	if ctx.KeyType() != nil {
		if ktype := ctx.KeyType().Accept(v); ktype != nil {
			field.KeyType, _ = ktype.(string)
		}
	}

	if ctx.Type_() != nil {
		if xtype := ctx.Type_().Accept(v); xtype != nil {
			field.ValType, _ = xtype.(string)
		}
	}

	if ctx.MapName() != nil {
		if name := ctx.MapName().Accept(v); name != nil {
			field.Name, _ = name.(string)
		}
	}

	return field
}

func (v *ProtoVisitor) VisitEnumDef(ctx *g4.EnumDefContext) any {
	enum := &EnumNode{
		Document: v.getHiddenTokensToLeft(ctx, ProtocolComment),
	}

	if ctx.EnumName() != nil {
		if name := ctx.EnumName().Accept(v); name != nil {
			enum.Name, _ = name.(string)
		}
	}

	if ctx.EnumBody() != nil {
		if body := ctx.EnumBody().Accept(v); body != nil {
			enum.Fields, _ = body.([]*EnumField)
		}
	}

	return enum
}

func (v *ProtoVisitor) VisitEnumName(ctx *g4.EnumNameContext) any {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitEnumBody(ctx *g4.EnumBodyContext) any {
	fields := make([]*EnumField, 0)

	for _, each := range ctx.AllEnumElement() {
		if field, _ := each.Accept(v).(*EnumField); field != nil {
			fields = append(fields, field)
		}
	}

	return fields
}

func (v *ProtoVisitor) VisitEnumElement(ctx *g4.EnumElementContext) any {
	if ctx.EnumField() == nil {
		return nil
	}

	return ctx.EnumField().Accept(v)
}

func (v *ProtoVisitor) VisitEnumField(ctx *g4.EnumFieldContext) any {
	field := &EnumField{
		Comment: v.getHiddenTokensToRight(ctx, ProtocolComment),
	}

	if ctx.Ident() != nil {
		if ident := ctx.Ident().Accept(v); ident != nil {
			field.Name, _ = ident.(string)
		}
	}

	if ctx.IntLit() != nil {
		if value := ctx.IntLit().Accept(v); value != nil {
			ivalue, _ := value.(string)
			field.Value, _ = strconv.Atoi(ivalue)
		}
	}

	return field
}

func (v *ProtoVisitor) VisitServiceDef(ctx *g4.ServiceDefContext) any {
	service := &ServiceNode{
		Document: v.getHiddenTokensToLeft(ctx, ProtocolComment),
	}

	if ctx.ServiceName() != nil {
		if name := ctx.ServiceName().Accept(v); name != nil {
			service.Name, _ = name.(string)
		}
	}

	for _, each := range ctx.AllServiceElement() {
		if method := each.Accept(v); method != nil {
			imethod, _ := method.(*ServiceMethod)
			service.Methods = append(service.Methods, imethod)
		}
	}

	return service
}

func (v *ProtoVisitor) VisitServiceName(ctx *g4.ServiceNameContext) any {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitServiceElement(ctx *g4.ServiceElementContext) any {
	if ctx.Rpc() == nil {
		return nil
	}

	return ctx.Rpc().Accept(v)
}

func (v *ProtoVisitor) VisitRpc(ctx *g4.RpcContext) any {
	method := &ServiceMethod{
		Document: v.getHiddenTokensToLeft(ctx, ProtocolComment),
		Options:  make(map[string]string),
	}

	if ctx.RpcName() != nil {
		if name := ctx.RpcName().Accept(v); name != nil {
			method.Name, _ = name.(string)
		}
	}

	for idx, each := range ctx.AllMessageType() {
		mtype := each.Accept(v)
		if mtype == nil {
			continue
		}
		if idx == 0 {
			method.ReqName, _ = mtype.(string)
		} else {
			method.ResName, _ = mtype.(string)
		}
	}

	for _, each := range ctx.AllOptionStatement() {
		if option, _ := each.Accept(v).(*OptionConst); option != nil {
			method.Options[option.Name] = option.Constant
		}
	}

	return method
}

func (v *ProtoVisitor) VisitRpcName(ctx *g4.RpcNameContext) any {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitMessageType(ctx *g4.MessageTypeContext) any {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitStrLit(ctx *g4.StrLitContext) any {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitIdent(ctx *g4.IdentContext) any {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitIntLit(ctx *g4.IntLitContext) any {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitOptionName(ctx *g4.OptionNameContext) any {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitConstant(ctx *g4.ConstantContext) any {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitKeyType(ctx *g4.KeyTypeContext) any {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitType_(ctx *g4.Type_Context) any {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitFieldName(ctx *g4.FieldNameContext) any {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitMapName(ctx *g4.MapNameContext) any {
	return ctx.GetText()
}

func (v *ProtoVisitor) getHiddenTokensToLeft(ctx ParseTopContext, channel int) string {
	stream, _ := ctx.GetParser().GetTokenStream().(*antlr.CommonTokenStream)
	tokens := stream.GetHiddenTokensToLeft(ctx.GetStart().GetTokenIndex(), channel)
	text := ""
	for _, each := range tokens {
		if each != nil && each.GetText() == "" {
			continue
		}
		text = strings.TrimPrefix(each.GetText(), "/*") // nolint:goconst
		text = strings.TrimPrefix(text, "//")           // nolint:goconst
		text = strings.TrimPrefix(text, " ")            // nolint:goconst
		text = strings.TrimSuffix(text, "*/")           // nolint:goconst
	}

	return text
}

func (v *ProtoVisitor) getHiddenTokensToRight(ctx ParseTopContext, channel int) string {
	stream, _ := ctx.GetParser().GetTokenStream().(*antlr.CommonTokenStream)
	tokens := stream.GetHiddenTokensToRight(ctx.GetStop().GetTokenIndex(), channel)
	text := ""
	for _, each := range tokens {
		if each != nil && each.GetText() == "" {
			continue
		}
		text = strings.TrimPrefix(each.GetText(), "/*")
		text = strings.TrimPrefix(text, "//")
		text = strings.TrimPrefix(text, " ")
		text = strings.TrimSuffix(text, "*/")

		break
	}

	return text
}

func isBasicType(btype string) bool {
	for _, t := range BasicTypes {
		if t == btype {
			return true
		}
	}

	return false
}
