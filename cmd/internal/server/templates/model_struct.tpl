// {{.Name}} {{.Document}}
type {{.Name}} struct {
  {{- range $field := $.Fields}}
    {{$field.Name}} {{$field.Type}} {{$field.Tag}} // {{$field.Comment}}
  {{- end}}
  {{- range $field := $.MapFields}}
    {{$field.Name}} map[{{$field.KeyType}}]{{$field.ValType}} {{$field.Tag}} // {{$field.Comment}}
  {{- end}}
}

