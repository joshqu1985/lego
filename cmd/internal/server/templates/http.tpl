package httpserver

import (
	"github.com/joshqu1985/lego/transport/rest"

  "{{$.ServerName}}/internal/service"
  "{{$.ServerName}}/internal/model"
)

// New{{$.Name}}Api 创建{{$.Name}}Api
func New{{$.Name}}Api(svc *service.Service) {{$.Name}}Api {
	return {{$.Name}}Api{
		svc,
	}
}

// {{$.Name}}Api
type {{$.Name}}Api struct {
	*service.Service
}

{{- range $method := $.Methods}}
// @Summary		  {{$method.Document}}
// @Description {{$method.Document}}
// @Accept		application/json
// @Tags		  {{$.ServerName}}/{{$method.ServiceName}}
{{if eq  $method.Cmd "Get" -}}
// @Param		  data	query     model.{{$method.ReqName}}				true	"data"
{{- else -}}
// @Param		  data	body      model.{{$method.ReqName}}				true	"data"
{{- end}}
// @Success		200		{object}	model.Response{data=model.{{$method.ResName}}}	"200 success"
// @failure		500		{object}	model.Response{data=string}	""
// @Router		{{$method.Uri}} [{{$method.Cmd}}]
// @Security	Bearer
func (this *{{$.Name}}Api) {{$method.Name}}(ctx *rest.Context) *rest.JSONResponse {
  var args model.{{$method.ReqName}}
	{{if eq  $method.Cmd "Get" -}}
  if err := ctx.ShouldBindQuery(&args); err != nil {
	{{- else -}}
  if err := ctx.ShouldBindJSON(&args); err != nil {
  {{- end}}
		return rest.ErrorResponse(400000, err.Error())
  }

  resp, err := this.{{$.Name}}.{{$method.Name}}(ctx, &args)
  if err != nil {
		return rest.ErrorResponse(500000, err.Error())
  }
	return rest.SuccessResponse(resp)
}
{{- end}}
