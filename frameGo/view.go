package frameGo

type View struct {
	selector string
	html     string
}

func NewView(selector, html string) View { return View{selector, html} }
