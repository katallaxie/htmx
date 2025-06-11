package tooltips

import htmx "github.com/katallaxie/htmx"

// Tip represents a tooltip with a message.
func Tip(text string) htmx.Node {
	return htmx.Attribute("data-tip", text)
}

// TooltipProps represents the properties for a tooltip component.
type TooltipProps struct {
	// Open indicates whether the tooltip is open or not.
	Open bool
	// Tip is the text to be displayed in the tooltip.
	Tip string
	// ClassNames contains the class names for the tooltip component.
	htmx.ClassNames
}

// Tooltip creates a tooltip component with the specified properties and children.
func Tooltip(p TooltipProps, children ...htmx.Node) htmx.Node {
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

// TooltipPrimary ...
func TooltipPrimary(p TooltipProps, children ...htmx.Node) htmx.Node {
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

// TooltipSecondary ...
func TooltipSecondary(p TooltipProps, children ...htmx.Node) htmx.Node {
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

// TooltipSuccess ...
func TooltipSuccess(p TooltipProps, children ...htmx.Node) htmx.Node {
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

// TooltipWarning ...
func TooltipWarning(p TooltipProps, children ...htmx.Node) htmx.Node {
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

// TooltipInfo ...
func TooltipInfo(p TooltipProps, children ...htmx.Node) htmx.Node {
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

// TooltipError ...
func TooltipError(p TooltipProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"tooltip":       true,
				"tooltip-error": true,
			},
			p.ClassNames,
		),
		Tip(p.Tip),
		htmx.Group(children...),
	)
}

// TooltipAccent ...
func TooltipAccent(p TooltipProps, children ...htmx.Node) htmx.Node {
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
