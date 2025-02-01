package pkg

import (
	"bytes"
	"text/template"
)

func Print(key, tpl string, data interface{}) string {
	tmpl, err := template.New(key).Parse(tpl)
	if err != nil {
		panic(err)
	}

	buffer := bytes.Buffer{}
	if err := tmpl.Execute(&buffer, data); err != nil {
		panic(err)
	}
	return buffer.String()
}
