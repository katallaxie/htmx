//nolint:revive,mnd
package main

import (
	"net/http"
	"time"

	"github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/imports"
	"github.com/katallaxie/htmx/imports/cache"
	"github.com/katallaxie/htmx/imports/jsdeliver"
	"github.com/katallaxie/htmx/tables"
	"github.com/katallaxie/pkg/conv"
)

const defaultTimeout = 3 * time.Second

type TableData struct {
	ID    int
	Name  string
	Email string
}

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
			tables.Table(
				tables.Props{
					Pagination: tables.Pagination(
						tables.PaginationProps{
							Limit: 10,
							Total: 3,
							URL:   "/",
							ClassNames: htmx.ClassNames{
								"m-4": true,
							},
						},
						tables.Prev(
							tables.PaginationProps{},
						),
						tables.Select(
							tables.PaginationProps{
								Limits: tables.DefaultLimits,
							},
						),
						tables.Next(
							tables.PaginationProps{},
						),
					),
				},
				tables.Columns[*TableData]{
					tables.ColumnDef[*TableData]{
						ID: "id",
						Header: func(p tables.Props) htmx.Node {
							return htmx.Th(
								htmx.Text("ID"),
							)
						},
						Cell: func(p tables.Props, row *TableData) htmx.Node {
							return htmx.Td(
								htmx.Text(
									conv.String(row.ID),
								),
							)
						},
					},
					tables.ColumnDef[*TableData]{
						ID: "name",
						Header: func(p tables.Props) htmx.Node {
							return htmx.Th(
								htmx.Text("Name"),
							)
						},
						Cell: func(p tables.Props, row *TableData) htmx.Node {
							return htmx.Td(
								htmx.Text(
									row.Name,
								),
							)
						},
					},
					tables.ColumnDef[*TableData]{
						ID: "email",
						Header: func(p tables.Props) htmx.Node {
							return htmx.Th(
								htmx.Text("Email"),
							)
						},
						Cell: func(p tables.Props, row *TableData) htmx.Node {
							return htmx.Td(
								htmx.Text(
									row.Email,
								),
							)
						},
					},
				},
				[]*TableData{
					{ID: 1, Name: "Alice", Email: "alice@example.com"},
					{ID: 2, Name: "Bob", Email: "bob@example.com"},
					{ID: 3, Name: "Charlie", Email: "charlie@example.com"},
				},
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
