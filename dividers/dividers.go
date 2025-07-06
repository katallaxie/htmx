package dividers

import htmx "github.com/katallaxie/htmx"

// Props is a struct that contains the props of a divider.
type Props struct {
	ClassNames htmx.ClassNames
}

// Divider is a struct that contains the props of a divider.
func Divider(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Neutral is a struct that contains the props of a neutral divider.
func Neutral(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider":         true,
				"divider-neutral": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Primary is a struct that contains the props of a primary divider.
func Primary(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider":         true,
				"divider-primary": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Secondary is a struct that contains the props of a secondary divider.
func Secondary(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider":           true,
				"divider-secondary": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Success is a struct that contains the props of a success divider.
func Success(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider":         true,
				"divider-success": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Warning is a struct that contains the props of a warning divider.
func Warning(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider":         true,
				"divider-warning": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Info is a struct that contains the props of an info divider.
func Info(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider":      true,
				"divider-info": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Error is a struct that contains the props of an error divider.
func Error(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"divider":      true,
				"divider-info": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
