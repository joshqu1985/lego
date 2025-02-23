func (this *{{.Name}}) ToModel() *model.{{.Name}} {
	m := &model.{{.Name}}{}
  {{- range $field := $.Fields}}
    {{if eq $field.Repeated true -}}
	    for _, item := range this.{{$field.Name}} {
        {{if eq $field.TypeObject true -}}
	      m.{{$field.Name}} = append(m.{{$field.Name}}, item.ToModel())
        {{- else -}}
	      m.{{$field.Name}} = append(m.{{$field.Name}}, item)
        {{- end}}
	    }
    {{- else -}}
      {{if eq $field.TypeObject true -}}
	    m.{{$field.Name}} = this.{{$field.Name}}.ToModel())
      {{- else -}}
      m.{{$field.Name}} = this.{{$field.Name}}
      {{- end}}
    {{- end}}
  {{- end}}
  return m
}

func (this *{{.Name}}) FromModel(data *model.{{.Name}}) *{{.Name}} {
  {{- range $field := $.Fields}}
    {{if eq $field.Repeated true -}}
	    for _, item := range data.{{$field.Name}} {
        {{if eq $field.TypeObject true -}}
	      this.{{$field.Name}} = append(this.{{$field.Name}}, (&{{$field.TypeName}}{}).FromModel(item))
        {{- else -}}
	      this.{{$field.Name}} = append(this.{{$field.Name}}, item)
        {{- end}}
	    }
    {{- else -}}
      {{if eq $field.TypeObject true -}}
      this.{{$field.Name}} = (&{{$field.TypeName}}{}).FromModel(data.{{$field.Name}})
      {{- else -}}
      this.{{$field.Name}} = data.{{$field.Name}}
      {{- end}}
    {{- end}}
  {{- end}}
  return this
}

