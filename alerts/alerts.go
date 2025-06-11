package alerts

import (
	htmx "github.com/katallaxie/htmx"
)

// AlertProps is the type of the props for the Alert component.
type Props struct {
	// ClassNames are the class names for the alert component.
	htmx.ClassNames
}

// Alert is a component that displays an alert.
func Alert(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"alert": true,
			},
			p.ClassNames,
		),
		htmx.Role("alert"),
		htmx.Group(children...),
	)
}

// Info is a component that displays an info alert.
func Info(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"alert":      true,
				"alert-info": true,
			},
			p.ClassNames,
		),
		htmx.Role("alert"),
		htmx.Group(children...),
	)
}

// Success is a component that displays a success alert.
func Success(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"alert":         true,
				"alert-success": true,
			},
			p.ClassNames,
		),
		htmx.Role("alert"),
		htmx.Group(children...),
	)
}

// Warning is a component that displays a warning alert.
func Warning(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"alert":         true,
				"alert-warning": true,
			},
			p.ClassNames,
		),
		htmx.Role("alert"),
		htmx.Group(children...),
	)
}

// Error is a component that displays an error alert.
func Error(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"alert":       true,
				"alert-error": true,
			},
			p.ClassNames,
		),
		htmx.Role("alert"),
		htmx.Group(children...),
	)
}

// Vertical is a component that displays a vertical alert.
func Vertical(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"alert":          true,
				"alert-vertical": true,
			},
			p.ClassNames,
		),
		htmx.Role("alert"),
		htmx.Group(children...),
	)
}

// Horizontal is a component that displays a horizontal alert.
func Horizontal(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"alert":            true,
				"alert-horizontal": true,
			},
			p.ClassNames,
		),
		htmx.Role("alert"),
		htmx.Group(children...),
	)
}
