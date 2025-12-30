package tables

import (
	"fmt"

	"github.com/katallaxie/htmx/buttons"
	"github.com/katallaxie/htmx/forms"
	"github.com/katallaxie/htmx/joins"
	"github.com/katallaxie/pkg/conv"
	"github.com/katallaxie/pkg/urlx"
	"github.com/katallaxie/pkg/utilx"

	htmx "github.com/katallaxie/htmx"
)

// DefaultLimits is a list of default limits.
var DefaultLimits = []int{5, 10, 25, 50}

// PaginationProps is a struct that contains the properties of a pagination.
type PaginationProps struct {
	// ID is the id of the table.
	ID string
	// Limit is the number of items to return.
	Limit int
	// Offset is the number of items to skip.
	Offset int
	// Target is the target of the pagination.
	Target string
	// Total is the total number of items.
	Total int
	// URL is the URL of the pagination.
	URL string
	// Limits is the list of limits.
	Limits []int

	htmx.ClassNames
}

// Pagination is a component that renders a pagination.
func Pagination(p PaginationProps, children ...htmx.Node) htmx.Node {
	return joins.Join(
		joins.Props{
			ClassNames: p.ClassNames,
		},
		htmx.Group(children...),
	)
}

// Prev is a component that renders a previous button.
func Prev(p PaginationProps) htmx.Node {
	return htmx.Form(
		htmx.Method("GET"),
		htmx.Action(urlx.MustRemoveQueryValues(p.URL, "offset", "limit")),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("offset"),
			htmx.Value(conv.String(p.Offset-p.Limit)),
		),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("limit"),
			htmx.Value(conv.String(p.Limit)),
		),
		htmx.HxBoost(true),
		buttons.Button(
			buttons.ButtonProps{
				ClassNames: htmx.Merge(
					htmx.ClassNames{
						"btn":            true,
						"input-bordered": true,
						"join-item":      true,
					},
					p.ClassNames,
				),
				Type: "submit",
			},
			htmx.If(p.Offset-p.Limit < 0, htmx.Disabled()),
			htmx.Text("Prev"),
		),
	)
}

// Next is a component that renders a next button.
func Next(p PaginationProps) htmx.Node {
	return htmx.Form(
		htmx.Method("GET"),
		htmx.Action(urlx.MustRemoveQueryValues(p.URL, "offset", "limit")),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("offset"),
			htmx.Value(conv.String(p.Offset+p.Limit)),
		),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("limit"),
			htmx.Value(conv.String(p.Limit)),
		),
		htmx.HxBoost(true),
		buttons.Button(
			buttons.ButtonProps{
				ClassNames: htmx.Merge(
					htmx.ClassNames{
						"join-item":      true,
						"btn":            true,
						"input-bordered": true,
					},
					p.ClassNames,
				),
				Type: "submit",
			},
			htmx.If(p.Offset+p.Limit > p.Total, htmx.Disabled()),
			htmx.Text("Next"),
		),
	)
}

// SearchProps are the properties of a search.
type SearchProps struct {
	// ClassNames is a struct that contains the class names of a search.
	ClassNames htmx.ClassNames
	// Placehholder is the placeholder of the search.
	Placeholder string
	// URL is the URL of the search.
	URL string
	// Name is the name of the search.
	Name string
	// Value is the value of the search.
	Value string
}

// Search is a component that renders a search.
func Search(props SearchProps, children ...htmx.Node) htmx.Node {
	return htmx.Form(
		htmx.Method("GET"),
		htmx.Action(urlx.MustRemoveQueryValues(props.URL, props.Name)),
		forms.TextInputBordered(
			forms.TextInputProps{
				ClassNames: htmx.Merge(
					props.ClassNames,
				),
				Name:        props.Name,
				Placeholder: props.Placeholder,
				Value:       props.Value,
			},
			htmx.Group(children...),
		),
		htmx.HxBoost(true),
	)
}

// Select is a component that renders a select.
func Select(p PaginationProps, children ...htmx.Node) htmx.Node {
	return htmx.Form(
		htmx.Method("GET"),
		htmx.Action(urlx.MustRemoveQueryValues(p.URL, "offset", "limit")),
		htmx.IfElse(utilx.NotEmpty(p.ID), htmx.HxTrigger(fmt.Sprintf("change from:#%s", p.ID)), htmx.HxTrigger("change from:#select-table-options")),
		htmx.Input(
			htmx.Type("hidden"),
			htmx.Name("offset"),
			htmx.Value(conv.String(p.Offset)),
		),
		htmx.HxBoost(true),
		joins.Item(
			joins.Props{
				ClassNames: p.ClassNames,
			},
		),
		forms.Select(
			forms.SelectProps{
				ClassNames: htmx.Merge(
					htmx.ClassNames{
						"join-item":      true,
						"input-bordered": true,
					},
					p.ClassNames,
				),
			},
			htmx.IfElse(utilx.NotEmpty(p.ID), htmx.ID(p.ID), htmx.ID("select-table-options")),
			htmx.Attribute("name", "limit"),
			htmx.Group(
				htmx.ForEach(
					p.Limits,
					func(limit int, _ int) htmx.Node {
						return forms.Option(
							forms.OptionProps{
								Selected: limit == p.Limit,
							},
							htmx.Text(conv.String(limit)),
							htmx.Value(conv.String(limit)),
						)
					},
				)...,
			),
		),
		htmx.Group(children...),
	)
}

// TableToolbarProps is a struct that contains the properties of a table toolbar.
type TableToolbarProps struct {
	ClassNames htmx.ClassNames
}

// TableToolbar is a component that renders a table toolbar.
func TableToolbar(p TableToolbarProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"table-toolbar": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// TablePaginationProps is a struct that contains the properties of a table pagination.
type TablePaginationProps struct {
	ClassNames htmx.ClassNames
}

// TablePagination is a component that renders a table pagination.
func TablePagination(p TablePaginationProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"flex":            true,
				"items-center":    true,
				"justify-between": true,
				"px-2":            true,
			},
			p.ClassNames,
		),
		htmx.Div(
			htmx.Merge(
				htmx.ClassNames{
					"flex":         true,
					"items-center": true,
					"space-x-6":    true,
					"lg:space-x-8": true,
				},
			),
			htmx.Group(children...),
		),
	)
}

// Row is a struct that contains the properties of a row.
type Row interface {
	comparable
}

// Props is a struct that contains the properties of a table.
type Props struct {
	// ID is the id of the table.
	ID string
	// Pagination is the pagination of the table.
	Pagination htmx.Node
	// Toolbar is the toolbar of the table.
	Toolbar htmx.Node

	htmx.ClassNames
}

// Columns returns a new column definition.
type Columns[R Row] []ColumnDef[R]

// ColumnDef returns a new column definition.
type ColumnDef[R Row] struct {
	// ID is the id of the column.
	ID string
	// AccessorKey is the accessor key of the column.
	AccessorKey string
	// Header is the header of the column.
	Header func(p Props) htmx.Node
	// Cell is the cell of the column.
	Cell func(p Props, row R) htmx.Node
	// EnableSorting is a flag to enable sorting.
	EnableSorting bool
	// EnableFiltering is a flag to enable filtering.
	EnableFiltering bool
}

// Table is a struct that contains the properties of a table.
func Table[S ~[]R, R Row](p Props, columns Columns[R], s S) htmx.Node {
	headers := []htmx.Node{}
	for _, column := range columns {
		headers = append(headers, column.Header(p))
	}

	rows := []htmx.Node{}
	for _, row := range s {
		cells := []htmx.Node{}
		for _, column := range columns {
			cells = append(cells, column.Cell(p, row))
		}
		rows = append(rows, htmx.Tr(cells...))
	}

	return htmx.Div(
		htmx.ID(p.ID),
		htmx.Merge(p.ClassNames),
		p.Toolbar,
		htmx.Div(
			htmx.Table(
				htmx.Merge(
					htmx.ClassNames{
						"table": true,
					},
					p.ClassNames,
				),
				htmx.THead(
					htmx.Tr(
						headers...,
					),
				),
				htmx.TBody(
					rows...,
				),
			),
		),
		p.Pagination,
	)
}
