package pkg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/iancoleman/strcase"

	"github.com/joshqu1985/lego/cmd/internal/pkg/g4"
)

func Parse(filename string) (*Tree, error) {
	if filename == "" {
		return nil, nil
	}

	file, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, err
	}

	stream := antlr.NewCommonTokenStream(g4.NewProtobuf3Lexer(file),
		antlr.TokenDefaultChannel)

	data := g4.NewProtobuf3Parser(stream).Proto().Accept(newProtoVisitor())

	tree := data.(*Tree)
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
		if isBasicType(field.Type) {
			if field.Repeated {
				field.Type = fmt.Sprintf("[]%s", field.Type)
			} else {
				field.Type = fmt.Sprintf("%s", field.Type)
			}
		} else {
			if field.Repeated {
				field.Type = fmt.Sprintf("[]*%s", field.Type)
			} else {
				field.Type = fmt.Sprintf("*%s", field.Type)
			}
		}
		vals := []string{}
		for _, tag := range tags {
			vals = append(vals, fmt.Sprintf("%s:\"%s\"", tag, strcase.ToSnake(field.Name)))
		}
		field.Tag = "`" + strings.Join(vals, " ") + "`"
	}

	for _, field := range item.MapFields {
		field.Name = strcase.ToCamel(field.Name)
		vals := []string{}
		for _, tag := range tags {
			vals = append(vals, fmt.Sprintf("%s:\"%s\"", tag, strcase.ToSnake(field.Name)))
		}
		field.Tag = "`" + strings.Join(vals, " ") + "`"
	}
}

type ProtoVisitor struct {
	g4.Protobuf3Visitor
	Tree Tree
}

func newProtoVisitor() *ProtoVisitor {
	return &ProtoVisitor{
		Tree: Tree{Options: map[string]string{}},
	}
}

func (v *ProtoVisitor) VisitProto(ctx *g4.ProtoContext) interface{} {
	for _, each := range ctx.AllImportStatement() {
		if statement := each.Accept(v); statement != nil {
			v.Tree.Imports = append(v.Tree.Imports, statement.(string))
		}
	}

	for _, each := range ctx.AllOptionStatement() {
		if option := each.Accept(v).(*OptionConst); option != nil {
			v.Tree.Options[option.Name] = option.Constant
		}
	}

	for _, each := range ctx.AllTopLevelDef() {
		each.Accept(v)
	}

	return &v.Tree
}

func (v *ProtoVisitor) VisitImportStatement(ctx *g4.ImportStatementContext) interface{} {
	if ctx.StrLit() == nil {
		return nil
	}
	return ctx.StrLit().Accept(v)
}

func (v *ProtoVisitor) VisitOptionStatement(ctx *g4.OptionStatementContext) interface{} {
	if ctx.OptionName() == nil || ctx.Constant() == nil {
		return nil
	}

	name := ctx.OptionName().Accept(v).(string)
	name = strings.TrimFunc(name, func(r rune) bool { return r == '(' || r == ')' })

	constant := ctx.Constant().Accept(v).(string)
	constant = strings.TrimFunc(constant, func(r rune) bool { return r == '"' })

	return &OptionConst{
		Name:     name,
		Constant: constant,
	}
}

func (v *ProtoVisitor) VisitTopLevelDef(ctx *g4.TopLevelDefContext) interface{} {
	if ctx.EnumDef() != nil {
		if enum := ctx.EnumDef().Accept(v); enum != nil {
			v.Tree.Enums = append(v.Tree.Enums, enum.(*EnumNode))
		}
	}

	if ctx.MessageDef() != nil {
		if message := ctx.MessageDef().Accept(v); message != nil {
			v.Tree.Structs = append(v.Tree.Structs, message.(*StructNode))
		}
	}

	if ctx.ServiceDef() != nil {
		if service := ctx.ServiceDef().Accept(v); service != nil {
			v.Tree.Services = append(v.Tree.Services, service.(*ServiceNode))
		}
	}

	return v
}

func (v *ProtoVisitor) VisitMessageDef(ctx *g4.MessageDefContext) interface{} {
	message := &StructNode{
		Document: v.getHiddenTokensToLeft(ctx, ProtocolComment),
		Options:  map[string]string{},
	}

	if ctx.MessageName() != nil {
		if name := ctx.MessageName().Accept(v); name != nil {
			message.Name = name.(string)
		}
	}

	if ctx.MessageBody() != nil {
		if body := ctx.MessageBody().Accept(v).(*StructNode); body != nil {
			for key, val := range body.Options {
				message.Options[key] = val
			}
			message.Fields = append(message.Fields, body.Fields...)
			message.MapFields = append(message.MapFields, body.MapFields...)
		}
	}

	return message
}

func (v *ProtoVisitor) VisitMessageName(ctx *g4.MessageNameContext) interface{} {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitMessageBody(ctx *g4.MessageBodyContext) interface{} {
	elements := &StructNode{
		Options: map[string]string{},
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

func (v *ProtoVisitor) VisitMessageElement(ctx *g4.MessageElementContext) interface{} {
	if ctx.MessageDef() != nil {
		if message := ctx.MessageDef().Accept(v); message != nil {
			v.Tree.Structs = append(v.Tree.Structs, message.(*StructNode))
		}
	}

	if ctx.EnumDef() != nil {
		if enum := ctx.EnumDef().Accept(v); enum != nil {
			v.Tree.Enums = append(v.Tree.Enums, enum.(*EnumNode))
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

func (v *ProtoVisitor) VisitField(ctx *g4.FieldContext) interface{} {
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
			field.Type = xtype.(string)
		}
	}

	if ctx.FieldName() != nil {
		if name := ctx.FieldName().Accept(v); name != nil {
			field.Name = name.(string)
		}
	}

	return field
}

func (v *ProtoVisitor) VisitMapField(ctx *g4.MapFieldContext) interface{} {
	field := &StructMap{
		Comment: v.getHiddenTokensToRight(ctx, ProtocolComment),
	}

	if ctx.KeyType() != nil {
		if ktype := ctx.KeyType().Accept(v); ktype != nil {
			field.KeyType = ktype.(string)
		}
	}

	if ctx.Type_() != nil {
		if xtype := ctx.Type_().Accept(v); xtype != nil {
			field.ValType = xtype.(string)
		}
	}

	if ctx.MapName() != nil {
		if name := ctx.MapName().Accept(v); name != nil {
			field.Name = name.(string)
		}
	}

	return field
}

func (v *ProtoVisitor) VisitEnumDef(ctx *g4.EnumDefContext) interface{} {
	enum := &EnumNode{
		Document: v.getHiddenTokensToLeft(ctx, ProtocolComment),
	}

	if ctx.EnumName() != nil {
		if name := ctx.EnumName().Accept(v); name != nil {
			enum.Name = name.(string)
		}
	}

	if ctx.EnumBody() != nil {
		if body := ctx.EnumBody().Accept(v); body != nil {
			enum.Fields = body.([]*EnumField)
		}
	}

	return enum
}

func (v *ProtoVisitor) VisitEnumName(ctx *g4.EnumNameContext) interface{} {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitEnumBody(ctx *g4.EnumBodyContext) interface{} {
	fields := []*EnumField{}

	for _, each := range ctx.AllEnumElement() {
		if field := each.Accept(v).(*EnumField); field != nil {
			fields = append(fields, field)
		}
	}

	return fields
}

func (v *ProtoVisitor) VisitEnumElement(ctx *g4.EnumElementContext) interface{} {
	if ctx.EnumField() == nil {
		return nil
	}
	return ctx.EnumField().Accept(v)
}

func (v *ProtoVisitor) VisitEnumField(ctx *g4.EnumFieldContext) interface{} {
	field := &EnumField{
		Comment: v.getHiddenTokensToRight(ctx, ProtocolComment),
	}

	if ctx.Ident() != nil {
		if ident := ctx.Ident().Accept(v); ident != nil {
			field.Name = ident.(string)
		}
	}

	if ctx.IntLit() != nil {
		if value := ctx.IntLit().Accept(v); value != nil {
			field.Value, _ = strconv.Atoi(value.(string))
		}
	}

	return field
}

func (v *ProtoVisitor) VisitServiceDef(ctx *g4.ServiceDefContext) interface{} {
	service := &ServiceNode{
		Document: v.getHiddenTokensToLeft(ctx, ProtocolComment),
	}

	if ctx.ServiceName() != nil {
		if name := ctx.ServiceName().Accept(v); name != nil {
			service.Name = name.(string)
		}
	}

	for _, each := range ctx.AllServiceElement() {
		if method := each.Accept(v); method != nil {
			service.Methods = append(service.Methods, method.(*ServiceMethod))
		}
	}

	return service
}

func (v *ProtoVisitor) VisitServiceName(ctx *g4.ServiceNameContext) interface{} {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitServiceElement(ctx *g4.ServiceElementContext) interface{} {
	if ctx.Rpc() == nil {
		return nil
	}
	return ctx.Rpc().Accept(v)
}

func (v *ProtoVisitor) VisitRpc(ctx *g4.RpcContext) interface{} {
	method := &ServiceMethod{
		Document: v.getHiddenTokensToLeft(ctx, ProtocolComment),
		Options:  map[string]string{},
	}

	if ctx.RpcName() != nil {
		if name := ctx.RpcName().Accept(v); name != nil {
			method.Name = name.(string)
		}
	}

	for idx, each := range ctx.AllMessageType() {
		if mtype := each.Accept(v); mtype != nil {
			if idx == 0 {
				method.ReqName = mtype.(string)
			} else {
				method.ResName = mtype.(string)
			}
		}
	}

	for _, each := range ctx.AllOptionStatement() {
		if option := each.Accept(v).(*OptionConst); option != nil {
			method.Options[option.Name] = option.Constant
		}
	}

	return method
}

func (v *ProtoVisitor) VisitRpcName(ctx *g4.RpcNameContext) interface{} {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitMessageType(ctx *g4.MessageTypeContext) interface{} {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitStrLit(ctx *g4.StrLitContext) interface{} {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitIdent(ctx *g4.IdentContext) interface{} {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitIntLit(ctx *g4.IntLitContext) interface{} {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitOptionName(ctx *g4.OptionNameContext) interface{} {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitConstant(ctx *g4.ConstantContext) interface{} {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitKeyType(ctx *g4.KeyTypeContext) interface{} {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitType_(ctx *g4.Type_Context) interface{} {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitFieldName(ctx *g4.FieldNameContext) interface{} {
	return ctx.GetText()
}

func (v *ProtoVisitor) VisitMapName(ctx *g4.MapNameContext) interface{} {
	return ctx.GetText()
}

const (
	ProtocolComment = 1
)

func (v *ProtoVisitor) getHiddenTokensToLeft(ctx ParseTopContext, channel int) (text string) {
	tokens := ctx.GetParser().GetTokenStream().(*antlr.CommonTokenStream).
		GetHiddenTokensToLeft(ctx.GetStart().GetTokenIndex(), channel)
	for _, each := range tokens {
		if each != nil && each.GetText() != "" {
			text = strings.TrimPrefix(each.GetText(), "/*")
			text = strings.TrimPrefix(text, "//")
			text = strings.TrimPrefix(text, " ")
			text = strings.TrimSuffix(text, "*/")
		}
	}
	return text
}

func (v *ProtoVisitor) getHiddenTokensToRight(ctx ParseTopContext, channel int) (text string) {
	tokens := ctx.GetParser().GetTokenStream().(*antlr.CommonTokenStream).
		GetHiddenTokensToRight(ctx.GetStop().GetTokenIndex(), channel)
	for _, each := range tokens {
		if each != nil && each.GetText() != "" {
			text = strings.TrimPrefix(each.GetText(), "/*")
			text = strings.TrimPrefix(text, "//")
			text = strings.TrimPrefix(text, " ")
			text = strings.TrimSuffix(text, "*/")
			break
		}
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
