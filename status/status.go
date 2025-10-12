package status

import htmx "github.com/katallaxie/htmx"

// Props is a struct that contains the properties of the Status component.
type Props struct {
	// AriaLabel is the aria-label attribute for the status element.
	AriaLabel string

	htmx.ClassNames
}

// Status is a component that renders a status indicator.
func Status(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status": true,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// Primary is a component that renders a primary status indicator.
func Primary(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":         true,
				"status-primary": true,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// Secondary is a component that renders a secondary status indicator.
func Secondary(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":           true,
				"status-secondary": true,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// Accent is a component that renders an accent status indicator.
func Accent(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":        true,
				"status-accent": true,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// Info is a component that renders an info status indicator.
func Info(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":      true,
				"status-info": true,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// Success is a component that renders a success status indicator.
func Success(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":         true,
				"status-success": true,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// Warning is a component that renders a warning status indicator.
func Warning(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":         true,
				"status-warning": true,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// Error is a component that renders an error status indicator.
func Error(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":       true,
				"status-error": true,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}

// Neutral is a component that renders a neutral status indicator.
func Neutral(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"status":         true,
				"status-neutral": true,
			},
			props.ClassNames,
		),
		htmx.Role("status"),
		htmx.Aria("label", props.AriaLabel),
		htmx.Group(children...),
	)
}
