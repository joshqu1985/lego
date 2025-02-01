package server

import (
	_ "embed"
	"fmt"
	"go/format"
	"path"
	// "path/filepath"
	"strings"

	"github.com/joshqu1985/lego/cmd/internal/pkg"
)

type ServiceTemplateVars struct {
	ServerName string // 服务名 (go mod name、registry name)
	Name       string // 模块名称
	Document   string // 模块文档
	Methods    []*MethodTemplateVars
}

type MethodTemplateVars struct {
	ServiceName string // 模块名
	Name        string // 方法名
	Document    string // 方法文档
	Cmd         string // 方法类型 rest POST / GET / DELETE
	Uri         string // 方法请求路径 rest
	ReqName     string // 入参自定义类型名
	ResName     string // 出参自定义类型名
}

var (
	//go:embed templates/http_init.tpl
	httpInitTpl string

	//go:embed templates/http.tpl
	httpTpl string

	//go:embed templates/rpc_init.tpl
	rpcInitTpl string

	//go:embed templates/rpc.tpl
	rpcTpl string

	//go:embed templates/model_head.tpl
	modelHeadTpl string

	//go:embed templates/model_enum.tpl
	modelEnumTpl string

	//go:embed templates/model_struct.tpl
	modelStructTpl string

	//go:embed templates/service_init.tpl
	serviceInitTpl string

	//go:embed templates/service.tpl
	serviceTpl string
)

func GenRestApis(name, src string, tree *pkg.Tree) ([]pkg.File, error) {
	files := []pkg.File{}
	for _, module := range tree.Services {
		codes, err := genRestModuleApis(name, src, module)
		if err != nil {
			return nil, err
		}
		files = append(files, codes...)
	}
	return files, nil
}

func GenRpcApis(name, src string, tree *pkg.Tree) ([]pkg.File, error) {
	files := []pkg.File{}
	for _, module := range tree.Services {
		codes, err := genRpcModuleApis(name, src, module)
		if err != nil {
			return nil, err
		}
		files = append(files, codes...)
	}
	return files, nil
}

func GenModels(name, src string, tree *pkg.Tree) ([]pkg.File, error) {
	vars := &ServiceTemplateVars{
		ServerName: name,
	}
	content := pkg.Print("model-head", modelHeadTpl, vars)

	for _, item := range tree.Enums {
		content = content + "\n" + pkg.Print("model-enum", modelEnumTpl, item)
	}
	for _, item := range tree.Structs {
		content = content + "\n" + pkg.Print("model-struct", modelStructTpl, item)
	}

	code, err := format.Source([]byte(content))
	if err != nil {
		return nil, err
	}
	return []pkg.File{
		{
			Name: fmt.Sprintf("/%s/internal/model/", name) +
				strings.Replace(path.Base(src), ".proto", ".gen.go", -1),
			Data:         code,
			ForceReplace: true,
		},
	}, nil
}

func GenServices(name, src string, tree *pkg.Tree) ([]pkg.File, error) {
	files := []pkg.File{}
	for _, module := range tree.Services {
		codes, err := genService(name, src, module)
		if err != nil {
			return nil, err
		}
		files = append(files, codes...)
	}
	return files, nil
}

func genRestModuleApis(serverName, src string, service *pkg.ServiceNode) ([]pkg.File, error) {
	vars := &ServiceTemplateVars{
		Name:       service.Name,
		Document:   service.Document,
		ServerName: serverName,
		Methods:    []*MethodTemplateVars{},
	}

	var ok bool
	for _, method := range service.Methods {
		mtd := &MethodTemplateVars{
			Name:        method.Name,
			Document:    method.Document,
			ServiceName: service.Name,
			ReqName:     method.ReqName,
			ResName:     method.ResName,
		}
		mtd.Cmd, ok = method.Options["method_cmd"]
		if !ok {
			panic("函数没有指定方法 method_cmd [GET POST PUT DELETE]")
		}
		mtd.Uri, ok = method.Options["method_uri"]
		if !ok {
			panic("函数没有指定路由 method_uri")
		}
		vars.Methods = append(vars.Methods, mtd)
	}

	code, err := format.Source([]byte(pkg.Print("http-apis", httpTpl, vars)))
	if err != nil {
		return nil, err
	}
	apiFile := pkg.File{
		Name: fmt.Sprintf("/%s/httpserver/", serverName) +
			strings.Replace(path.Base(src), ".proto", ".go", -1),
		Data: code,
	}

	value := pkg.Print("http-init-apis", httpInitTpl, vars)
	code, err = format.Source([]byte(pkg.Print("http-init-apis", httpInitTpl, vars)))
	if err != nil {
		fmt.Println("-------------------2", value, err)
		return nil, err
	}
	apiInitFile := pkg.File{
		Name: fmt.Sprintf("/%s/httpserver/httpserver.go", serverName),
		Data: code,
	}
	return []pkg.File{apiFile, apiInitFile}, nil
}

func genRpcModuleApis(serverName, src string, service *pkg.ServiceNode) ([]pkg.File, error) {
	vars := &ServiceTemplateVars{
		Name:       service.Name,
		Document:   service.Document,
		ServerName: serverName,
		Methods:    []*MethodTemplateVars{},
	}
	for _, method := range service.Methods {
		vars.Methods = append(vars.Methods, &MethodTemplateVars{
			Name:        method.Name,
			Document:    method.Document,
			ServiceName: service.Name,
			ReqName:     method.ReqName,
			ResName:     method.ResName,
		})
	}

	value := pkg.Print("rpc-apis", rpcTpl, vars)
	code, err := format.Source([]byte(pkg.Print("rpc-apis", rpcTpl, vars)))
	if err != nil {
		fmt.Println("-------------------2", value, err)
		return nil, err
	}
	apiFile := pkg.File{
		Name: fmt.Sprintf("/%s/rpcserver/", serverName) +
			strings.Replace(path.Base(src), ".proto", ".go", -1),
		Data: code,
	}

	code, err = format.Source([]byte(pkg.Print("rpc-init-apis", rpcInitTpl, vars)))
	if err != nil {
		fmt.Println("-------------------1", err)
		return nil, err
	}
	apiInitFile := pkg.File{
		Name: fmt.Sprintf("/%s/rpcserver/rpcserver.go", serverName),
		Data: code,
	}
	return []pkg.File{apiFile, apiInitFile}, nil
}

func genService(serverName, src string, service *pkg.ServiceNode) ([]pkg.File, error) {
	vars := &ServiceTemplateVars{
		Document:   service.Document,
		ServerName: serverName,
		Name:       service.Name,
		Methods:    []*MethodTemplateVars{},
	}
	for _, method := range service.Methods {
		vars.Methods = append(vars.Methods, &MethodTemplateVars{
			Document: method.Document,
			Name:     method.Name,
			ReqName:  method.ReqName,
			ResName:  method.ResName,
		})
	}

	code, err := format.Source([]byte(pkg.Print("service", serviceTpl, vars)))
	if err != nil {
		return nil, err
	}
	serviceFile := pkg.File{
		Name: fmt.Sprintf("/%s/internal/service/", serverName) +
			strings.Replace(path.Base(src), ".proto", ".go", -1),
		Data: code,
	}

	code, err = format.Source([]byte(pkg.Print("service-init", serviceInitTpl, vars)))
	if err != nil {
		return nil, err
	}
	serviceInitFile := pkg.File{
		Name: fmt.Sprintf("/%s/internal/service/service.go", serverName),
		Data: code,
	}
	return []pkg.File{serviceFile, serviceInitFile}, nil
}
