/* {{.Document}} */
type {{.Name}} struct {
  {{- range $field := $.Fields}}
    {{$field.Name}} {{$field.Type}} {{$field.Tag}} // {{$field.Comment}}
  {{- end}}
  {{- range $field := $.MapFields}}
    {{$field.Name}} map[{{$field.KeyType}}]{{$field.ValType}} {{$field.Tag}} // {{$field.Comment}}
  {{- end}}
}

{{- if .IsForm}}
func (this *{{.Name}}) Encode() map[string]string {
	params := map[string]string{}
  {{- range $field := $.Fields}}
	  {{if eq  $field.Type "string" -}}
    params["{{$field.Name}}"] = this.{{$field.Name}}
    {{- else -}}
    params["{{$field.Name}}"] = fmt.Sprintf("%v", this.{{$field.Name}})
    {{- end}}
  {{- end}}
  return params
}
{{- end}}
