package pages

import (
	"stack/App/components"
	"stack/frameGo"
)

type Login struct {
}

func (login *Login) View() (views []frameGo.View) {
	l := components.NewLogin()
	view := frameGo.NewView("body", l.Render())
	views = append(views, view)
	return
}
