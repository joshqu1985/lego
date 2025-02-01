// {{$.Name}}Rpc {{$.Document}}
type {{$.Name}}Rpc struct {
	client {{$.Name}}Client
}

func New{{$.Name}}Rpc(target string, n naming.Naming) (*{{$.Name}}Rpc, error) {
	c, err := rpc.NewClient(target, rpc.WithNaming(n))
	if err != nil {
		return nil, err
	}
	return &{{$.Name}}Rpc{client: New{{$.Name}}Client(c.Conn())}, nil
}

{{- range $method := $.Methods}}
// {{$method.Name}} {{$method.Document}}
func (this *{{$.Name}}Rpc) {{$method.Name}}(ctx context.Context, in *{{$method.ReqName}})(*{{$method.ResName}}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := this.client.{{$method.Name}}(ctx, in)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
{{- end}}
