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
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// BadgePrimary is a function that returns a primary indicator item.
func BadgePrimary(props ItemProps, children ...htmx.Node) htmx.Node {
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

// BadgeSecondary is a function that returns a secondary indicator item.
func BadgeSecondary(props ItemProps, children ...htmx.Node) htmx.Node {
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

// StatusSuccess is a function that returns a success indicator item.
func StatusSuccess(props ItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"indicator-item": true,
				"status":         true,
				"status-success": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// StatusError is a function that returns an error indicator item.
func StatusError(props ItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"indicator-item": true,
				"status":         true,
				"status-error":   true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
