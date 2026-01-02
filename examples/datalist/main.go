package main

import (
	"net/http"
	"time"

	"github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/htmx/imports"
	"github.com/katallaxie/htmx/imports/cache"
	"github.com/katallaxie/htmx/imports/jsdeliver"
	"github.com/katallaxie/htmx/loading"
	"github.com/katallaxie/pkg/conv"
)

const defaultTimeout = 3 * time.Second

func Page() htmx.Node {
	return htmx.HTML5(
		htmx.HTML5Props{
			Title: "HTMX Page",
			Head: []htmx.Node{
				htmx.Link(
					htmx.Href("https://cdn.jsdelivr.net/npm/daisyui@5"),
					htmx.Rel("stylesheet"),
					htmx.Type("text/css"),
				),
				htmx.Script(
					htmx.Src("https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"),
				),
				htmx.Imports(
					htmx.ImportsProp{
						Resolver: cache.New(jsdeliver.New()),
						Pkgs: []imports.ExactPackage{
							{
								Name:    "htmx.org",
								Version: "2.0.4",
							},
						},
						Requires: []imports.Require{
							{
								File: "dist/htmx.esm.js",
							},
						},
					},
				),
				htmx.Script(
					htmx.Type("module"),
					htmx.Raw(`import htmx from "htmx.org";`),
				),
			},
		},
		htmx.Body(
			htmx.Div(
				forms.Datalist(
					forms.DatalistProps{
						ID:          "users-datalist",
						Target:      "#users-datalist",
						URL:         "/users",
						Name:        "user",
						Placeholder: "Search for a user...",
						Indicator:   "#user-search-indicator",
					},
				),
				loading.Spinner(
					loading.SpinnerProps{},
					htmx.ID("user-search-indicator"),
				),
			),
		),
	)
}

func hello(w http.ResponseWriter, _ *http.Request) {
	_ = Page().Render(w)
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/users", func(w http.ResponseWriter, _ *http.Request) {
		users := []string{"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi"}
		resp := htmx.Fragment(
			htmx.ForEach(users, func(u string, idx int) htmx.Node {
				return htmx.Option(
					htmx.Value(conv.String(idx)),
					htmx.Text(u),
				)
			})...,
		)

		_ = resp.Render(w)
	})

	server := &http.Server{
		Addr:              ":3000",
		ReadHeaderTimeout: defaultTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
