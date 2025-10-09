package main

import (
	"net/http"
	"time"

	"github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/drawers"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/htmx/imports"
	"github.com/katallaxie/htmx/imports/cache"
	"github.com/katallaxie/htmx/imports/jsdeliver"
	"github.com/katallaxie/htmx/loading"
	"github.com/katallaxie/htmx/menus"
	"github.com/katallaxie/htmx/navbars"
	"github.com/katallaxie/htmx/tailwind"
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
				forms.FormControl(
					forms.FormControlProps{},
					forms.Datalist(
						forms.DatalistProps{
							ID:          "workflows",
							Name:        "workflow_id",
							Target:      "#workflows",
							Placeholder: "Search a workflow ...",
							URL:         "/workflows",
						},
					),
					loading.Spinner(
						loading.SpinnerProps{
							ClassNames: htmx.ClassNames{
								"htmx-indicator": true,
								tailwind.M2:      true,
							},
						},
					),
				),
				htmx.Div(
					forms.LabelInput(
						forms.LabelProps{},
						forms.LabelText(
							forms.LabelProps{},
							htmx.Text("Select a date"),
						),
						forms.DateInput(
							forms.DateInputProps{},
						),
					),
				),
				forms.LabelSelect(
					forms.LabelProps{},
					htmx.Span(
						htmx.Class("label"),
						htmx.Text("Select an option"),
					),
					forms.Select(
						forms.SelectProps{},
						forms.Option(
							forms.OptionProps{
								Value:    "option1",
								Selected: true,
							},
							htmx.Text("Option 1"),
						),
						forms.Option(
							forms.OptionProps{
								Value: "option2",
							},
							htmx.Text("Option 2"),
						),
					),
				),

				htmx.Div(
					htmx.ClassNames{
						"flex":           true,
						"flex-col":       true,
						"items-center":   true,
						"justify-center": true,
						"flex-1":         true,
						"p-4":            true,
					},
					forms.Fieldset(
						forms.FieldsetProps{
							ClassNames: htmx.ClassNames{
								"bg-base-200":     true,
								"border-base-300": true,
								"rounded-box":     true,
								"border":          true,
								"p-4":             true,
								"m-4":             true,
								"w-300":           true,
							},
						},
						forms.Legend(
							forms.LegendProps{},
							htmx.Text("Page Title"),
						),
						forms.TextInput(
							forms.TextInputProps{
								Required: true,
								ClassNames: htmx.ClassNames{
									"input-bordered": true,
									"w-full":         true,
								},
							},
						),
						forms.Label(
							forms.LabelProps{},
							htmx.Text("This is an example page using HTMX components."),
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
	http.HandleFunc("/workflows", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		htmx.Fragment(
			htmx.Option(
				htmx.Value("1"),
				htmx.Text("Workflow 1"),
			),
			htmx.Option(
				htmx.Value("2"),
				htmx.Text("Workflow 2"),
			),
		).Render(w)
	})

	server := &http.Server{
		Addr:              ":3000",
		ReadHeaderTimeout: defaultTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
