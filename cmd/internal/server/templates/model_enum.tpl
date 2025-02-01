// {{.Name}} {{.Document}}
type {{.Name}} int32

const (
    {{- range $field := $.Fields}}
      {{$.Name}}_{{$field.Name}} {{$.Name}} = {{$field.Value}} // {{$field.Comment}}
    {{- end}}
)

