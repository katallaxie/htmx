package main

import (
	"net/http"
	"time"

	"github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/htmx/imports"
	"github.com/katallaxie/htmx/imports/cache"
	"github.com/katallaxie/htmx/imports/jsdeliver"
	"github.com/katallaxie/htmx/mockups"
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
								Version: "4.0.1",
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
			mockups.Window(
				mockups.WindowProps{
					ClassNames: htmx.ClassNames{
						"m-4": true,
					},
				},
				htmx.Div(
					htmx.ClassNames{
						"p-4": true,
					},
					forms.RangeSuccess(
						forms.RangeProps{
							Min:        "0",
							Max:        "100",
							Value:      "40",
							ClassNames: htmx.ClassNames{},
						},
					),
					forms.RangeNeutral(
						forms.RangeProps{
							Min:        "0",
							Max:        "100",
							Value:      "40",
							ClassNames: htmx.ClassNames{},
						},
					),
					forms.RangeVertical(
						forms.RangeProps{
							Min:        "0",
							Max:        "100",
							Value:      "40",
							ClassNames: htmx.ClassNames{},
						},
					),
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

	server := &http.Server{
		Addr:              ":3000",
		ReadHeaderTimeout: defaultTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
