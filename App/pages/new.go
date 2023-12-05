package pages

import (
	"stack/App/components"
	"stack/frameGo"
)

type New struct {
}

func (n *New) View() []frameGo.View {
	principal := components.NewPrincipal("Clash⚡", "Latest clashs", "5")
	principalRender := frameGo.NewView("#clash-title", principal.Render())
	clashRender := frameGo.NewView("#clash", "")
	var views []frameGo.View
	views = append(views, principalRender, clashRender)
	return views
}
