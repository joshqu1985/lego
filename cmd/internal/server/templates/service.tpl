package service

import (
  "context"

	faker "github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"

	"{{$.ServerName}}/internal/model"
  // "{{$.ServerName}}/internal/repo"
)

type I{{$.Name}} interface {
{{- range $method := $.Methods}}
  // {{$method.Document}}
  {{$method.Name}}(ctx context.Context, req *model.{{$method.ReqName}}) (*model.{{$method.ResName}}, error)
{{- end}}
}

type {{$.Name}}Impl struct {
}

func New{{$.Name}}() I{{$.Name}} {
  s := &{{$.Name}}Impl{}
  return s
}

{{- range $method := $.Methods}}
// {{$method.Name}} {{$method.Document}}
func (this *{{$.Name}}Impl) {{$method.Name}}(ctx context.Context, req *model.{{$method.ReqName}}) (*model.{{$method.ResName}}, error) {
  resp := &model.{{$method.ResName}}{}
  // TODO MOCK
	return resp, faker.FakeData(resp, options.WithRandomMapAndSliceMaxSize(3), options.WithRandomStringLength(7))
}
{{end}}
