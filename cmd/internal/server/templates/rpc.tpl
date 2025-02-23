package rpcserver

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

  "{{$.ServerName}}/internal/service"
	pb "{{$.ServerName}}/rpcserver/protocols"
)

// New{{$.Name}}Api 创建{{$.Name}}Api
func New{{$.Name}}Api(svc *service.Service) *{{$.Name}}Api {
	return &{{$.Name}}Api{
		Service: svc,
	}
}

// {{$.Name}}Api
type {{$.Name}}Api struct {
	*service.Service
	pb.Unimplemented{{$.Name}}Server
}

{{- range $method := $.Methods}}
// {{$method.Name}} {{$method.Document}}
func (this *{{$.Name}}Api) {{$method.Name}}(ctx context.Context, args *pb.{{$method.ReqName}}) (*pb.{{$method.ResName}}, error) {
  resp := &pb.{{$method.ResName}}{}
  data, err := this.{{$.Name}}.{{$method.Name}}(ctx, args.ToModel())
  if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
  }
	return resp.FromModel(data), nil
}
{{- end}}
