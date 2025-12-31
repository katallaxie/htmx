package main

import (
	"net/http"
	"time"

	"github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/avatars"
	"github.com/katallaxie/htmx/drawers"
	"github.com/katallaxie/htmx/dropdowns"
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
			drawers.Drawer(
				drawers.Props{
					Open: true,
				},
				drawers.Toggle(
					drawers.ToggleProps{
						ID: "app-drawer",
					},
				),
				drawers.Content(
					drawers.ContentProps{},
					navbars.Navbar(
						navbars.Props{
							ClassNames: htmx.ClassNames{
								"w-full":      true,
								"bg-base-300": true,
							},
						},
						htmx.Div(
							htmx.ClassNames{
								"px-4": true,
							},
							htmx.Text("Title"),
						),
					),
				),
				drawers.Side(
					drawers.SideProps{
						ClassNames: htmx.ClassNames{
							"w-80":            true,
							"bg-base-200":     true,
							"min-h-full":      true,
							"flex":            true,
							"flex-col":        true,
							"justify-between": true,
						},
					},
					menus.Menu(
						menus.Props{
							ClassNames: htmx.ClassNames{
								"bg-base-200": true,
								"rounded-box": true,
								"w-80":        true,
								"p-4":         true,
							},
						},
						menus.Title(
							menus.TitleProps{},
							htmx.Text("Menu"),
						),
						menus.Item(
							menus.ItemProps{},
							htmx.A(
								htmx.Href("#"),
								htmx.Text("Item 1"),
							),
						),
						menus.Title(
							menus.TitleProps{},
							htmx.Text("Submenu"),
						),
						menus.Item(
							menus.ItemProps{},
							menus.Collapsible(
								menus.CollapsibleProps{},
								menus.CollapsibleSummary(
									menus.CollapsibleSummaryProps{},
									htmx.Text("Parent"),
								),
								menus.Menu(
									menus.Props{
										ClassNames: htmx.ClassNames{
											"menu": false,
										},
									},
									menus.Item(
										menus.ItemProps{},
										htmx.A(
											htmx.Href("#"),
											htmx.Text("Submenu 1"),
										),
									),
									menus.Item(
										menus.ItemProps{},
										htmx.A(
											htmx.Href("#"),
											htmx.Text("Submenu 2"),
										),
									),
								),
							),
						),
						menus.Title(
							menus.TitleProps{},
							htmx.Text("Menu"),
						),
						menus.Item(
							menus.ItemProps{},
							htmx.A(
								htmx.Href("#"),
								htmx.Text("Item 2"),
							),
						),
					),
					dropdowns.Dropdown(
						dropdowns.Props{
							ClassNames: htmx.ClassNames{
								"w-full": true,
								"p-4":    true,
								"py-8":   true,
							},
						},
						dropdowns.Button(
							dropdowns.ButtonProps{
								ClassNames: htmx.ClassNames{
									"w-full":        true,
									"flex":          true,
									"justify-start": true,
								},
							},
							avatars.RoundedSmall(
								avatars.Props{},
								htmx.Img(
									htmx.Src("https://img.daisyui.com/images/profile/demo/batperson@192.webp"),
								),
							),
							htmx.Text("Indy Jones"),
						),
						dropdowns.MenuItems(
							dropdowns.MenuItemsProps{},
							htmx.Li(
								htmx.A(
									htmx.Href("#"),
									htmx.Text("Option 1"),
								),
							),
							htmx.Li(
								htmx.A(
									htmx.Href("#"),
									htmx.Text("Option 2"),
								),
							),
						),
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
