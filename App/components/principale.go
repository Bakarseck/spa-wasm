package components

import "stack/frameGo"

type Principal struct {
	Mode  string
	Title string
	With  string
}

func (p *Principal) Render() string {
	component :=
		`
	<div style="width: {{.With}}em;" id="clash-label">{{.Mode}}</div>
    <h2 id="titre2">{{.Title}}</h2>
	`
	return frameGo.ParseComponent(component, p)
}

func NewPrincipal(mode, title, width string) Principal {
	return Principal{
        Mode:  mode,
        Title: title,
        With:  width,
    }
}
