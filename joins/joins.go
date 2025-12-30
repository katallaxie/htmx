package joins

import htmx "github.com/katallaxie/htmx"

// Props is a struct that contains the properties of a join.
type Props struct {
	htmx.ClassNames
}

// Join is a function that returns a join.
func Join(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"join": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Vertical is a function that returns a vertical join.
func Vertical(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"join":          true,
				"join-vertical": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Horizontal is a function that returns a horizontal join.
func Horizontal(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"join":            true,
				"join-horizontal": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Item is a function that returns a join item.
func Item(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"join-item": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
