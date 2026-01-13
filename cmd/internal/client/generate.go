package client

import (
	"fmt"
	"go/format"
	"path"
	"path/filepath"
	"strings"

	"github.com/codeskyblue/go-sh"

	_ "embed"

	"github.com/joshqu1985/lego/cmd/internal/pkg"
)

type (
	ServiceTemplateVars struct {
		PackageName string
		Name        string // 模块名称
		Document    string // 模块文档
		Methods     []*MethodTemplateVars
	}

	MethodTemplateVars struct {
		Name     string
		Document string
		Cmd      string
		Uri      string
		ReqName  string
		ResName  string
		IsForm   bool
	}
)

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
		Name: fmt.Sprintf("/%s/", packageName) + //nolint:goconst
			strings.ReplaceAll(path.Base(src), pkg.FileSuffixProto, "_model"+pkg.FileSuffixGenGo),
		Data:         code,
		ForceReplace: true,
	}, nil
}

func genRestRequest(packageName, src string, tree *pkg.Tree) (pkg.File, error) {
	vars := &ServiceTemplateVars{
		PackageName: packageName,
	}
	content := pkg.Print("rest-method-head", restMethodHeadTpl, vars)

	stus := make(map[string]*pkg.StructNode, 0)
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
			strings.ReplaceAll(path.Base(src), pkg.FileSuffixProto, pkg.FileSuffixGenGo),
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
		content = content + "\n" + genRpcService(packageName, service)
	}

	code, err := format.Source([]byte(content))
	if err != nil {
		return pkg.File{}, err
	}

	return pkg.File{
		Name: fmt.Sprintf("/%s/", packageName) +
			strings.ReplaceAll(path.Base(src), pkg.FileSuffixProto, pkg.FileSuffixGenGo),
		Data:         code,
		ForceReplace: true,
	}, nil
}

func genRestEnum(e *pkg.EnumNode) string {
	return pkg.Print("rest-model-enum", restModelEnumTpl, e)
}

func genRestStruct(s *pkg.StructNode) string {
	var structVars struct {
		*pkg.StructNode
		IsForm bool
	}
	structVars.StructNode = s

	if tags := s.Options["message_tag"]; strings.Contains(tags, pkg.StringForm) {
		structVars.IsForm = true
	}

	return pkg.Print("rest-model-struct", restModelStructTpl, structVars)
}

func genRestService(stus map[string]*pkg.StructNode, service *pkg.ServiceNode) string {
	vars := &ServiceTemplateVars{
		Document: service.Document,
		Name:     service.Name,
		Methods:  make([]*MethodTemplateVars, 0),
	}

	var mok bool
	for _, method := range service.Methods {
		methodVar := &MethodTemplateVars{
			Document: method.Document,
			Name:     method.Name,
			ReqName:  method.ReqName,
			ResName:  method.ResName,
		}
		methodVar.Cmd, mok = method.Options["method_cmd"]
		if !mok {
			continue
		}
		methodVar.Uri, mok = method.Options["method_uri"]
		if !mok {
			continue
		}

		request, sok := stus[method.ReqName]
		if !sok {
			continue
		}
		if tags := request.Options["message_tag"]; strings.Contains(tags, pkg.StringForm) {
			methodVar.IsForm = true
		}
		vars.Methods = append(vars.Methods, methodVar)
	}

	return pkg.Print("rest-method-body", restMethodBodyTpl, vars)
}

func genRpcService(packageName string, service *pkg.ServiceNode) string {
	vars := &ServiceTemplateVars{
		PackageName: packageName,
		Document:    service.Document,
		Name:        service.Name,
		Methods:     make([]*MethodTemplateVars, 0),
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
