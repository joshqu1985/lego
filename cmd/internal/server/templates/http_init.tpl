package httpserver

import (
	"github.com/joshqu1985/lego/transport/rest"

  "{{$.ServerName}}/internal/service"
)

// New 创建HTTPServer
func New(svc *service.Service) *HTTPServer {
	return &HTTPServer{
		{{$.Name}}: New{{$.Name}}Api(svc),
	}
}

// HTTPServer
type HTTPServer struct {
	{{$.Name}} {{$.Name}}Api
}

// 路由
func (this *HTTPServer) Routers(r *rest.Router) {
	{
	{{- range $method := $.Methods}}
	// {{$method.Document}}
	r.{{$method.Cmd}}("{{$method.Uri}}", rest.ResponseWrapper(this.{{$.Name}}.{{$method.Name}}))
	{{- end}}
	}
}

