package pkg

import (
	"bytes"
	"text/template"
)

func Print(key, tpl string, data any) string {
	tmpl, err := template.New(key).Parse(tpl)
	if err != nil {
		panic(err)
	}

	buffer := bytes.Buffer{}
	if xerr := tmpl.Execute(&buffer, data); xerr != nil {
		panic(xerr)
	}

	return buffer.String()
}
