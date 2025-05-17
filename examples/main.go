package main

import (
	"net/http"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
)

func Demo() htmx.Node {
	return htmx.HTML5(
		htmx.HTML5Props{
			Title:    "Demo",
			Language: "en",
			Head: []htmx.Node{
				htmx.Script(
					htmx.Src("https://unpkg.com/fiber-reload@0.9.0/reload.js"),
					htmx.Type("text/javascript"),
				),
			},
		},
		htmx.Body(
			buttons.Button(buttons.ButtonProps{}, htmx.Text("Demo")),
		),
	)
}

func hello(w http.ResponseWriter, req *http.Request) {
	Demo().Render(w)
}

func main() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":3000", nil)
}
