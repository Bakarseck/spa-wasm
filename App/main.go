package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"stack/App/pages"
	"stack/frameGo"
	"syscall/js"
)

func main() {
	ch := make(chan chan struct{})
	fmt.Println("WebAssembly from Golang")

	router := &frameGo.Router{}

	routes := map[string]pages.Pages{
		"/":         &pages.Home{},
		"/new":      &pages.New{},
		"/login":    &pages.Login{},
		"/register": &pages.Register{},
	}

	for route, page := range routes {
		router.AddRoute(route, page.View())
	}

	home := js.Global().Get("location").Get("href").String()
	router.RouteResolver(home)

	elements := js.Global().Get("document").Call("querySelectorAll", "a")
	for i := 0; i < elements.Length(); i++ {
		element := elements.Index(i)
		element.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			event := args[0]
			event.Call("preventDefault")
			router.RouteResolver(element.Get("href").String())
			return nil
		}))
	}

	element := js.Global().Get("document").Call("querySelector", "#submitLogin")
	element.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		args[0].Call("preventDefault")

		login := js.Global().Get("document").Call("querySelector", "#loginUsername").Get("value").String()
		password := js.Global().Get("document").Call("querySelector", "#loginPassword").Get("value").String()

		data := map[string]string{
			"login":    login,
			"password": password,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err.Error())
		}

		resp, err := http.Post("http://localhost:8083/login", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println(err.Error())
		}
		defer resp.Body.Close()

		content, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(string(content))

		return nil
	}))
	<-ch
}
