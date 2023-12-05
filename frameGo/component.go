package frameGo

import (
	"bytes"
	"html/template"
)

type Component interface {
	Render() string
}

func ParseComponent(component string, data any) string {
	tmpl, err := template.New("component").Parse(component)
	if err != nil {
		return err.Error()
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}
