
// {{.SingularName}} {{.Comment}}
type {{.SingularName}} struct {
	{{- range $col := $.Columns}}
	{{$col.Name}} {{$col.DataType}} {{$col.Tag}} // {{$col.Comment}}
	{{- end}}
}

func (this *{{.SingularName}}) TableName() string {
  return "{{.Name}}"
}
