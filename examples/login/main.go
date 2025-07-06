package main

import (
	"net/http"
	"time"

	"github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/cards"
	"github.com/katallaxie/htmx/dividers"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/htmx/imports"
	"github.com/katallaxie/htmx/imports/cache"
	"github.com/katallaxie/htmx/imports/jsdeliver"
	"github.com/katallaxie/htmx/links"
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
			htmx.Section(
				htmx.Merge(
					htmx.ClassNames{
						"bg-gray-50":       true,
						"dark:bg-gray-900": true,
					},
				),
			),
			htmx.Div(
				htmx.Merge(
					htmx.ClassNames{
						"flex":           true,
						"flex-col":       true,
						"items-center":   true,
						"justify-center": true,
						"px-6":           true,
						"py-8":           true,
						"mx-auto":        true,
						"md:h-screen":    true,
						"lg:py-0":        true,
					},
				),
				cards.CardBorder(
					cards.CardProps{
						ClassNames: htmx.ClassNames{
							"w-full":   true,
							"max-w-md": true,
						},
					},
					cards.Body(
						cards.BodyProps{},
						cards.Title(
							cards.TitleProps{},
							htmx.Text("Sign in to your account"),
						),
						htmx.Div(
							htmx.ClassNames{},
							links.Button(
								links.Props{
									ClassNames: htmx.ClassNames{
										"w-full": true,
									},
									Href: "/login/entraid",
								},
								htmx.Text("Login on Microsoft Entra ID"),
							),
						),
						htmx.Div(
							htmx.ClassNames{},
							links.Button(
								links.Props{
									ClassNames: htmx.ClassNames{
										"w-full": true,
									},
									Href: "/login/github",
								},
								htmx.Text("Login on GitHub"),
							),
						),
						dividers.Divider(
							dividers.Props{},
							htmx.Text("OR"),
						),
						htmx.Form(
							htmx.HxPost("/login"),
							forms.Fieldset(
								forms.FieldsetProps{},
								forms.Legend(
									forms.LegendProps{},
									htmx.Text("Login with your credentials"),
								),
								forms.Label(
									forms.LabelProps{},
									htmx.Text("Username"),
								),
								forms.TextInput(
									forms.TextInputProps{
										Name:        "username",
										Placeholder: "indy@jones.com",
									},
								),
								forms.Label(
									forms.LabelProps{},
									htmx.Text("Password"),
								),
								forms.TextInput(
									forms.TextInputProps{
										Name: "password",
										Type: "password",
									},
								),
							),
							buttons.Primary(
								buttons.ButtonProps{
									ClassNames: htmx.ClassNames{
										"w-full": true,
										"my-4":   true,
									},
								},
								htmx.Text("Login"),
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
