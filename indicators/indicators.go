package indicators

import htmx "github.com/katallaxie/htmx"

// Props is a struct that contains the properties of an indicator.
type Props struct {
	htmx.ClassNames
}

// Indicator is a function that returns an indicator.
func Indicator(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"indicator": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ItemProps is a struct that contains the properties of an indicator item.
type ItemProps struct {
	htmx.ClassNames
}

// Item is a function that returns an indicator item.
func Item(props ItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"indicator-item": true,
				"badge":          true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ItemPrimary is a function that returns a primary indicator item.
func ItemPrimary(props ItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"indicator-item": true,
				"badge":          true,
				"badge-primary":  true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ItemSecondary is a function that returns a secondary indicator item.
func ItemSecondary(props ItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"indicator-item":  true,
				"badge":           true,
				"badge-secondary": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
