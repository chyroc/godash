package internal

import (
	"bytes"
	"text/template"
)

type TemplateData map[string]interface{}

func BuildTemplate(tpl string, data interface{}) string {
	t, err := template.New("").Parse(tpl)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
