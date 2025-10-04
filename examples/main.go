package main

import (
	"net/http"
	"time"

	"github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/drawers"
	"github.com/katallaxie/htmx/imports"
	"github.com/katallaxie/htmx/imports/cache"
	"github.com/katallaxie/htmx/imports/jsdeliver"
	"github.com/katallaxie/htmx/menus"
	"github.com/katallaxie/htmx/navbars"
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
		drawers.Drawer(
			drawers.Props{},
			drawers.DrawerToggle(
				drawers.ToggleProps{
					ID: "my-drawer",
				},
			),
			drawers.DrawerContent(
				drawers.DrawerContentProps{
					ClassNames: htmx.ClassNames{
						"flex":     true,
						"flex-col": true,
					},
				},
				htmx.Div(),
				navbars.Navbar(
					navbars.Props{
						ClassNames: htmx.ClassNames{
							"bg-base-300": true,
							"w-full":      true,
						},
					},
					htmx.Div(
						htmx.ClassNames{
							"flex-1": true,
							"px-2":   true,
							"mx-2":   true,
						},
						htmx.Text("My App"),
					),
					htmx.Div(
						htmx.ClassNames{
							"flex-none": true,
							"block":     true,
						},
						menus.MenuHorizontal(
							menus.Props{},
							menus.Item(
								menus.ItemProps{},
								htmx.A(
									htmx.Href("#"),
									htmx.Text("Item 1"),
								),
							),
							menus.Item(
								menus.ItemProps{},
								htmx.A(
									htmx.Href("#"),
									htmx.Text("Item 2"),
								),
							),
						),
					),
				),

				htmx.Text("Hello, World!"),
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
