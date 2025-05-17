package main

import (
	"net/http"
	"time"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
)

const defaultTimeout = 3 * time.Second

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

func hello(w http.ResponseWriter, _ *http.Request) {
	_ = Demo().Render(w)
}

func main() {
	http.HandleFunc("/", hello)

	server := &http.Server{
		Addr:              ":3000",
		ReadHeaderTimeout: defaultTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
