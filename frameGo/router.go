package frameGo

import (
	"regexp"
	"syscall/js"
)

type Route struct {
	path  string
	views []View
}

type Router struct {
	routes []Route
}

func (r *Router) AddRoute(p string, v []View) {
	myRoutes := Route{p, v}
	r.routes = append(r.routes, myRoutes)
}

func (r *Router) RouteResolver(URL string) {
	var element js.Value
	match := MatchPath(URL)
	for i, route := range r.routes {
		if route.path == match {
			for _, view := range route.views {
				element = js.Global().Get("document").Call("querySelector", view.selector)
				element.Set("innerHTML", view.html)
				js.Global().Get("history").Call("pushState", js.Null(), "", route.path)
			}
			break
		} else if i == len(r.routes)-1 {
			element = js.Global().Get("document").Call("querySelector", "body")
			element.Set("innerHTML", "<h1>Error 404: Page Not Found")
		}
	}
}

func MatchPath(url string) string {
	pathRegex := regexp.MustCompile(`http://[^/]+(/.*)?`)
	match := pathRegex.FindStringSubmatch(url)
	if match != nil {
		return match[1]
	}
	return ""
}
