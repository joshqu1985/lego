package client

import (
	_ "embed"
	"fmt"
	"go/format"
	"path"
	"path/filepath"
	"strings"

	"github.com/codeskyblue/go-sh"

	"github.com/joshqu1985/lego/cmd/internal/pkg"
)

type ServiceTemplateVars struct {
	PackageName string
	Name        string // 模块名称
	Document    string // 模块文档
	Methods     []*MethodTemplateVars
}

type MethodTemplateVars struct {
	Name     string // 方法名
	Document string // 方法文档
	Cmd      string // 方法类型 rest Post / Get / Delete
	Uri      string // 方法请求路径 rest
	IsForm   bool   // 入参是否是form表单 rest
	ReqName  string // 入参自定义类型名
	ResName  string // 出参自定义类型名
}

var (
	//go:embed templates/rest_method_head.tpl
	restMethodHeadTpl string

	//go:embed templates/rest_method_body.tpl
	restMethodBodyTpl string

	//go:embed templates/rest_model_head.tpl
	restModelHeadTpl string

	//go:embed templates/rest_model_enum.tpl
	restModelEnumTpl string

	//go:embed templates/rest_model_struct.tpl
	restModelStructTpl string

	//go:embed templates/rpc_method_head.tpl
	rpcMethodHeadTpl string

	//go:embed templates/rpc_method_body.tpl
	rpcMethodBodyTpl string
)

func genRestModel(packageName, src string, tree *pkg.Tree) (pkg.File, error) {
	vars := &ServiceTemplateVars{
		PackageName: packageName,
	}
	content := pkg.Print("rest-model-head", restModelHeadTpl, vars)

	for _, item := range tree.Enums {
		content = content + "\n" + genRestEnum(item)
	}

	for _, item := range tree.Structs {
		content = content + "\n" + genRestStruct(item)
	}

	code, err := format.Source([]byte(content))
	if err != nil {
		return pkg.File{}, err
	}

	return pkg.File{
		Name: fmt.Sprintf("/%s/", packageName) +
			strings.Replace(path.Base(src), ".proto", "_model.gen.go", -1),
		Data:         code,
		ForceReplace: true,
	}, nil
}

func genRestRequest(packageName, src string, tree *pkg.Tree) (pkg.File, error) {
	vars := &ServiceTemplateVars{
		PackageName: packageName,
	}
	content := pkg.Print("rest-method-head", restMethodHeadTpl, vars)

	stus := map[string]*pkg.StructNode{}
	for _, stu := range tree.Structs {
		stus[stu.Name] = stu
	}
	for _, service := range tree.Services {
		content = content + "\n" + genRestService(stus, service)
	}

	code, err := format.Source([]byte(content))
	if err != nil {
		return pkg.File{}, err
	}

	return pkg.File{
		Name: fmt.Sprintf("/%s/", packageName) +
			strings.Replace(path.Base(src), ".proto", ".gen.go", -1),
		Data:         code,
		ForceReplace: true,
	}, nil
}

func genRpcProtoFile(packageName, src, dst string) error {
	args := []any{
		"--proto_path=" + filepath.Dir(src),
		"--go_out=" + dst,
		"--go-grpc_out=" + dst,
		fmt.Sprintf("--go_opt=M%s=./%s", filepath.Base(src), packageName),
		fmt.Sprintf("--go-grpc_opt=M%s=./%s", filepath.Base(src), packageName),
		src,
	}
	if err := sh.Command("protoc", args...).Run(); err != nil {
		return err
	}
	return nil
}

func genRpcRequest(packageName, src string, tree *pkg.Tree) (pkg.File, error) {
	vars := &ServiceTemplateVars{
		PackageName: packageName,
	}
	content := pkg.Print("rpc-method-head", rpcMethodHeadTpl, vars)

	for _, service := range tree.Services {
		content = content + "\n" + genRpcService(service)
	}

	code, err := format.Source([]byte(content))
	if err != nil {
		return pkg.File{}, err
	}

	return pkg.File{
		Name: fmt.Sprintf("/%s/", packageName) +
			strings.Replace(path.Base(src), ".proto", ".gen.go", -1),
		Data:         code,
		ForceReplace: true,
	}, nil
}

func genRestEnum(e *pkg.EnumNode) string {
	return pkg.Print("rest-model-enum", restModelEnumTpl, e)
}

func genRestStruct(s *pkg.StructNode) string {
	var structVars struct {
		IsForm bool
		*pkg.StructNode
	}
	structVars.StructNode = s

	if tags, _ := s.Options["message_tag"]; strings.Contains(tags, "form") {
		structVars.IsForm = true
	}
	return pkg.Print("rest-model-struct", restModelStructTpl, structVars)
}

func genRestService(stus map[string]*pkg.StructNode, service *pkg.ServiceNode) string {
	vars := &ServiceTemplateVars{
		Document: service.Document,
		Name:     service.Name,
		Methods:  []*MethodTemplateVars{},
	}

	var ok bool
	for _, method := range service.Methods {
		methodVar := &MethodTemplateVars{
			Document: method.Document,
			Name:     method.Name,
			ReqName:  method.ReqName,
			ResName:  method.ResName,
		}
		methodVar.Cmd, ok = method.Options["method_cmd"]
		if !ok {
			continue
		}
		methodVar.Uri, ok = method.Options["method_uri"]
		if !ok {
			continue
		}

		request, ok := stus[method.ReqName]
		if !ok {
			continue
		}
		if tags, _ := request.Options["message_tag"]; strings.Contains(tags, "form") {
			methodVar.IsForm = true
		}
		vars.Methods = append(vars.Methods, methodVar)
	}
	return pkg.Print("rest-method-body", restMethodBodyTpl, vars)
}

func genRpcService(service *pkg.ServiceNode) string {
	vars := &ServiceTemplateVars{
		Document: service.Document,
		Name:     service.Name,
		Methods:  []*MethodTemplateVars{},
	}

	var ok bool
	for _, method := range service.Methods {
		methodVar := &MethodTemplateVars{
			Document: method.Document,
			Name:     method.Name,
			ReqName:  method.ReqName,
			ResName:  method.ResName,
		}
		methodVar.Cmd, ok = method.Options["method_cmd"]
		if ok {
			continue
		}
		methodVar.Uri, ok = method.Options["method_uri"]
		if ok {
			continue
		}
		vars.Methods = append(vars.Methods, methodVar)
	}
	return pkg.Print("rpc-method-body", rpcMethodBodyTpl, vars)
}
