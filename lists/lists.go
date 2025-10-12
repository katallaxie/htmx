package lists

import htmx "github.com/katallaxie/htmx"

// Props is a struct that contains the properties of the List component.
type Props struct {
	htmx.ClassNames
}

// List is a component that renders a list of items.
func List(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Ul(
		htmx.Merge(
			htmx.ClassNames{
				"list": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// TitleProps is a component that renders a header for a list.
type TitleProps struct {
	htmx.ClassNames
}

// Title is a component that renders a header for a list.
func Title(props TitleProps, children ...htmx.Node) htmx.Node {
	return htmx.Li(
		htmx.Merge(
			htmx.ClassNames{
				"opacity-60":    true,
				"p-4":           true,
				"pb-2":          true,
				"text-xs":       true,
				"tracking-wide": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// RowProps is a component that renders a row in a list.
type RowProps struct {
	htmx.ClassNames
}

// Row is a component that renders a row in a list.
func Row(props RowProps, children ...htmx.Node) htmx.Node {
	return htmx.Li(
		htmx.Merge(
			htmx.ClassNames{
				"list-row": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
