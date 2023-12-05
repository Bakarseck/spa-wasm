package pages

import (
	"stack/App/components"
	"stack/frameGo"
)

type Home struct {
	clashs []components.Clash
}

func (h *Home) View() []frameGo.View {
	clash1 := components.NewClash(0, 10, 7, "Angular", "React")
	clash2 := components.NewClash(1, 120, 54, "Golang", "Rust")

	tabClash := []components.Clash{clash1, clash2}
	h.clashs = append(h.clashs, tabClash...)
	render := ""

	for _, v := range h.clashs {
		render += v.Render()
	}

	principal := components.NewPrincipal("Clash⚡", "Last created clash", "5")

	clashRender := frameGo.NewView("#clash", render)
	principalRender := frameGo.NewView("#clash-title", principal.Render())

	var views []frameGo.View
	views = append(views, clashRender, principalRender)
	return views
}
