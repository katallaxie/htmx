package tooltips

import htmx "github.com/katallaxie/htmx"

// Tip represents a tooltip with a message.
func Tip(text string) htmx.Node {
	return htmx.Attribute("data-tip", text)
}

// Props represents the properties for a tooltip component.
type Props struct {
	// Open indicates whether the tooltip is open or not.
	Open bool
	// Tip is the text to be displayed in the tooltip.
	Tip string
	// ClassNames contains the class names for the tooltip component.
	htmx.ClassNames
}

// Tooltip creates a tooltip component with the specified properties and children.
func Tooltip(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":      true,
				"tooltip-open": p.Open,
			},
			p.ClassNames,
		),
		Tip(p.Tip),
		htmx.Group(children...),
	)
}

// Primary creates a primary tooltip component with the specified properties and children.
func Primary(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":         true,
				"tooltip-primary": true,
				"tooltip-open":    p.Open,
			},
			p.ClassNames,
		),
		Tip(p.Tip),
		htmx.Group(children...),
	)
}

// Secondary creates a secondary tooltip component with the specified properties and children.
func Secondary(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":           true,
				"tooltip-secondary": true,
				"tooltip-open":      p.Open,
			},
			p.ClassNames,
		),
		Tip(p.Tip),
		htmx.Group(children...),
	)
}

// Success creates a success tooltip with the specified properties and children.
func Success(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":         true,
				"tooltip-success": true,
				"tooltip-open":    p.Open,
			},
			p.ClassNames,
		),
		Tip(p.Tip),
		htmx.Group(children...),
	)
}

// Warning is a struct that contains the props of a warning tooltip.
func Warning(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":         true,
				"tooltip-warning": true,
				"tooltip-open":    p.Open,
			},
			p.ClassNames,
		),
		Tip(p.Tip),
		htmx.Group(children...),
	)
}

// Info is a struct that contains the props of an info tooltip.
func Info(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":      true,
				"tooltip-info": true,
				"tooltip-open": p.Open,
			},
			p.ClassNames,
		),
		Tip(p.Tip),
		htmx.Group(children...),
	)
}

// Error creates a tooltip component with error styling.
func Error(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":       true,
				"tooltip-error": true,
				"tooltip-open":  p.Open,
			},
			p.ClassNames,
		),
		Tip(p.Tip),
		htmx.Group(children...),
	)
}

// Accent is a struct that contains the props of an accent tooltip.
func Accent(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":        true,
				"tooltip-accent": true,
				"tooltip-open":   p.Open,
			},
			p.ClassNames,
		),
		Tip(p.Tip),
		htmx.Group(children...),
	)
}
