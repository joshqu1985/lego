// {{$.Name}}Rest {{$.Document}}
type {{$.Name}}Rest struct {
	client *rest.Client
}

func New{{$.Name}}Rest(target string, n naming.Naming) (*{{$.Name}}Rest, error) {
	c, err := rest.NewClient(target, rest.WithNaming(n))
	if err != nil {
		return nil, err
	}
	return &{{$.Name}}Rest{client: c}, nil
}

{{- range $method := $.Methods}}
// {{$method.Name}} {{$method.Document}}
func (this *{{$.Name}}Rest) {{$method.Name}}(ctx context.Context, req *{{$method.ReqName}})(*{{$method.ResName}}, error) {
	response, err := this.client.Request().
  {{- if .IsForm}}
		SetQueryParams(req.Encode()).
	{{- else -}}
		SetBody(req).
  {{- end}}
	{{if eq  $method.Cmd "Get" -}}
		Get("{{$method.Uri}}")
	{{- else -}}
		Post("{{$method.Uri}}")
  {{- end}}
	if err != nil {
		return nil, err
	}
	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("invalid http code %d", response.StatusCode())
	}

  var resp struct {
    Code int                  `json:"code"`
    Msg  string               `json:"msg"`
    Data *{{$method.ResName}} `json:"data"`
  }
	if err := json.Unmarshal(response.Body(), &resp); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("%d-%s", resp.Code, resp.Msg)
	}
	return resp.Data, nil
}
{{- end}}
