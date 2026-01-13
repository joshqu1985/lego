package server

import (
	"fmt"
	"go/format"
	"path"
	"path/filepath"
	"strings"

	"github.com/codeskyblue/go-sh"
	"github.com/golang/glog"
	"github.com/iancoleman/strcase"
	"github.com/samber/lo"

	_ "embed"

	"github.com/joshqu1985/lego/cmd/internal/pkg"
)

type (
	ServiceTemplateVars struct {
		ServerName string // 服务名 (go mod name、registry name)
		Name       string // 模块名称
		Document   string // 模块文档
		Methods    []*MethodTemplateVars
	}

	MethodTemplateVars struct {
		ServiceName string // 模块名
		Name        string // 方法名
		Document    string // 方法文档
		Cmd         string // 方法类型 rest POST / GET / DELETE
		Uri         string // 方法请求路径 rest
		ReqName     string // 入参自定义类型名
		ResName     string // 出参自定义类型名
	}

	Table struct {
		Name         string
		Comment      string
		SingularName string    `gorm:"-"`
		PluralName   string    `gorm:"-"`
		Columns      []*Column `gorm:"-"`
	}

	Column struct {
		Name     string
		DataType string
		Tag      string
		Comment  string
	}
)

var (
	//go:embed templates/http_main.tpl
	httpMainTpl string

	//go:embed templates/http_init.tpl
	httpInitTpl string

	//go:embed templates/http.tpl
	httpTpl string

	//go:embed templates/rpc_main.tpl
	rpcMainTpl string

	//go:embed templates/rpc_init.tpl
	rpcInitTpl string

	//go:embed templates/rpc.tpl
	rpcTpl string

	//go:embed templates/rpc_conv_head.tpl
	rpcConvHeadTpl string

	//go:embed templates/rpc_conv.tpl
	rpcConvTpl string

	//go:embed templates/config.tpl
	configTpl string

	//go:embed templates/model_head.tpl
	modelHeadTpl string

	//go:embed templates/model_enum.tpl
	modelEnumTpl string

	//go:embed templates/model_struct.tpl
	modelStructTpl string

	//go:embed templates/model_entity.tpl
	modelEntityTpl string

	//go:embed templates/service_init.tpl
	serviceInitTpl string

	//go:embed templates/service.tpl
	serviceTpl string

	//go:embed templates/repo_init.tpl
	repoInitTpl string

	MysqlTypes = map[string]string{
		"tinyint":           "int8",
		"tinyint unsigned":  "uint8",
		"smallint":          "int16",
		"smallint unsigned": "uint16",
		"integer":           "int64",
		"int":               "int",
		"int unsigned":      "uint",
		"bigint":            "int64",
		"bigint unsigned":   "uint64",
		"float":             "float32",
		"float unsigned":    "float32",
		"double":            "float64",
		"decimal":           "float64",
		"varchar":           "string",
		"char":              "string",
		"date":              "string",
		"time":              "string",
		"datetime":          "string",
		"timestamp":         "string",
		"json":              "datatypes.JSON",
		"text":              "[]byte",
		"tinytext":          "[]byte",
		"mediumtext":        "[]byte",
		"longtext":          "[]byte",
		"blob":              "[]byte",
		"tinyblob":          "[]byte",
		"mediumblob":        "[]byte",
		"longblob":          "[]byte",
		"enum":              "string",
	}
)

func GenRestMain(name string) (pkg.File, error) {
	vars := &ServiceTemplateVars{
		ServerName: name,
	}

	content := pkg.Print("http-main", httpMainTpl, vars)
	code, err := format.Source([]byte(content))
	if err != nil {
		return pkg.File{}, err
	}

	return pkg.File{
		Name: fmt.Sprintf("/%s/cmd/main.go", name), //nolint:goconst
		Data: code,
	}, nil
}

func GenRestApis(name, src string, tree *pkg.Tree) ([]pkg.File, error) {
	files := make([]pkg.File, 0)
	for _, module := range tree.Services {
		codes, err := restModuleApis(name, src, module)
		if err != nil {
			return nil, err
		}
		files = append(files, codes...)
	}

	return files, nil
}

func GenRpcMain(name string) (pkg.File, error) {
	vars := &ServiceTemplateVars{
		ServerName: name,
	}

	content := pkg.Print("rpc-main", rpcMainTpl, vars)
	code, err := format.Source([]byte(content))
	if err != nil {
		return pkg.File{}, err
	}

	return pkg.File{
		Name: fmt.Sprintf("/%s/cmd/main.go", name),
		Data: code,
	}, nil
}

func GenRpcApis(name, src string, tree *pkg.Tree) ([]pkg.File, error) {
	files := make([]pkg.File, 0)
	for _, module := range tree.Services {
		codes, err := rpcModuleApis(name, src, module)
		if err != nil {
			return nil, err
		}
		files = append(files, codes...)
	}

	return files, nil
}

func GenRpcProtos(name, src, dst string, tree *pkg.Tree) ([]pkg.File, error) {
	if err := grpcCodes(name, src, dst); err != nil {
		return nil, err
	}

	file, err := grpcConverter(name, src, tree.Structs)
	if err != nil {
		return nil, err
	}

	return []pkg.File{file}, nil
}

func GenConfig(name string) (pkg.File, error) {
	content := pkg.Print("config", configTpl, nil)
	code, err := format.Source([]byte(content))
	if err != nil {
		return pkg.File{}, err
	}

	return pkg.File{
		Name: fmt.Sprintf("/%s/internal/config/config.go", name),
		Data: code,
	}, nil
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
				strings.ReplaceAll(path.Base(src), pkg.FileSuffixProto, pkg.FileSuffixGenGo),
			Data:         code,
			ForceReplace: true,
		},
	}, nil
}

func GenServices(name, src string, tree *pkg.Tree) ([]pkg.File, error) {
	files := make([]pkg.File, 0)
	for _, module := range tree.Services {
		codes, err := serviceCode(name, src, module)
		if err != nil {
			return nil, err
		}
		files = append(files, codes...)
	}

	return files, nil
}

func GenRepository(name string) (pkg.File, error) {
	content := pkg.Print("repo-init", repoInitTpl, nil)
	code, err := format.Source([]byte(content))
	if err != nil {
		return pkg.File{}, err
	}

	return pkg.File{
		Name: fmt.Sprintf("/%s/internal/repo/repo.go", name),
		Data: code,
	}, nil
}

func GenModelEntities(name string, tables []*Table) ([]pkg.File, error) {
	files := make([]pkg.File, 0)
	for _, table := range tables {
		file, err := modelEntity(name, table)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

func restModuleApis(sname, src string, service *pkg.ServiceNode) ([]pkg.File, error) {
	vars := &ServiceTemplateVars{
		Name:       service.Name,
		Document:   service.Document,
		ServerName: sname,
		Methods:    make([]*MethodTemplateVars, 0),
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
		Name: fmt.Sprintf("/%s/httpserver/", sname) +
			strings.ReplaceAll(path.Base(src), pkg.FileSuffixProto, pkg.FileSuffixGo),
		Data: code,
	}

	code, err = format.Source([]byte(pkg.Print("http-init-apis", httpInitTpl, vars)))
	if err != nil {
		return nil, err
	}
	apiInitFile := pkg.File{
		Name: fmt.Sprintf("/%s/httpserver/httpserver.go", sname),
		Data: code,
	}

	return []pkg.File{apiFile, apiInitFile}, nil
}

func rpcModuleApis(sname, src string, service *pkg.ServiceNode) ([]pkg.File, error) {
	vars := &ServiceTemplateVars{
		Name:       service.Name,
		Document:   service.Document,
		ServerName: sname,
		Methods:    make([]*MethodTemplateVars, 0),
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

	code, err := format.Source([]byte(pkg.Print("rpc-apis", rpcTpl, vars)))
	if err != nil {
		return nil, err
	}
	apiFile := pkg.File{
		Name: fmt.Sprintf("/%s/rpcserver/", sname) +
			strings.ReplaceAll(path.Base(src), pkg.FileSuffixProto, pkg.FileSuffixGo),
		Data: code,
	}

	code, err = format.Source([]byte(pkg.Print("rpc-init-apis", rpcInitTpl, vars)))
	if err != nil {
		return nil, err
	}
	apiInitFile := pkg.File{
		Name: fmt.Sprintf("/%s/rpcserver/rpcserver.go", sname),
		Data: code,
	}

	return []pkg.File{apiFile, apiInitFile}, nil
}

func grpcCodes(name, src, dst string) error {
	path := fmt.Sprintf("./%s/%s/rpcserver/", dst, name)
	args := []any{
		"--proto_path=" + filepath.Dir(src),
		"--go_out=" + path,
		"--go-grpc_out=" + path,
		fmt.Sprintf("--go_opt=M%s=./protocols", filepath.Base(src)),
		fmt.Sprintf("--go-grpc_opt=M%s=./protocols", filepath.Base(src)),
		src,
	}
	if err := sh.Command("protoc", args...).Run(); err != nil {
		glog.Error(args, err)

		return err
	}

	return nil
}

func grpcConverter(name, src string, nodes []*pkg.StructNode) (pkg.File, error) {
	vars := &ServiceTemplateVars{
		ServerName: name,
	}
	content := pkg.Print("rpc-conv-head", rpcConvHeadTpl, vars)

	dict := lo.SliceToMap(nodes, func(node *pkg.StructNode) (string, struct{}) {
		return node.Name, struct{}{}
	})
	for _, node := range nodes {
		for _, field := range node.Fields {
			fieldType := strings.Trim(field.Type, "[]*")
			if _, ok := dict[fieldType]; !ok {
				continue
			}
			field.TypeObject, field.TypeName = true, fieldType
		}
	}

	for _, node := range nodes {
		content = content + "\n" + pkg.Print("rpc-conv", rpcConvTpl, node)
	}

	code, err := format.Source([]byte(content))
	if err != nil {
		return pkg.File{}, err
	}

	file := pkg.File{
		Name: fmt.Sprintf("/%s/rpcserver/protocols/", name) +
			strings.ReplaceAll(path.Base(src), pkg.FileSuffixProto, "_conv"+pkg.FileSuffixGenGo),
		Data:         code,
		ForceReplace: true,
	}

	return file, nil
}

func serviceCode(serverName, src string, service *pkg.ServiceNode) ([]pkg.File, error) {
	vars := &ServiceTemplateVars{
		Document:   service.Document,
		ServerName: serverName,
		Name:       service.Name,
		Methods:    make([]*MethodTemplateVars, 0),
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
			strings.ReplaceAll(path.Base(src), pkg.FileSuffixProto, pkg.FileSuffixGo),
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

func modelEntity(name string, table *Table) (pkg.File, error) {
	hasJson := false
	for _, column := range table.Columns {
		if column.DataType == "json" {
			hasJson = true
		}
		column.Name = strcase.ToCamel(column.Name)
		column.DataType = switchMysqlType(column.DataType)
		column.Tag = fmt.Sprintf("`json:%q`", strcase.ToSnake(column.Name))
	}
	table.SingularName = strcase.ToCamel(table.SingularName)

	content := `// Code generated by lego. DO NOT EDIT.

package entity
`
	if hasJson {
		content += `

import "gorm.io/datatypes"
    `
	}
	content += pkg.Print("model-entity", modelEntityTpl, table)

	code, err := format.Source([]byte(content))
	if err != nil {
		return pkg.File{}, err
	}

	return pkg.File{
		Name:         fmt.Sprintf("/%s/internal/model/entity/", name) + table.Name + ".gen.go",
		Data:         code,
		ForceReplace: true,
	}, nil
}

func switchMysqlType(dtype string) string {
	gotype, ok := MysqlTypes[dtype]
	if !ok {
		panic("unknown mysql type")
	}

	return gotype
}
