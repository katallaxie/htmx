package badges

import htmx "github.com/katallaxie/htmx"

// Props represents the properties for a badge element.
type Props struct {
	htmx.ClassNames // The class names for the badge element.
}

// Badge generates a badge element based on the provided properties.
func Badge(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"badge": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Neutral generates a neutral badge element based on the provided properties.
func Neutral(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"badge":         true,
				"badge-neutral": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Primary generates a primary badge element based on the provided properties.
func Primary(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"badge":         true,
				"badge-primary": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Secondary generates a secondary badge element based on the provided properties.
func Secondary(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"badge":           true,
				"badge-secondary": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Accent generates an accent badge element based on the provided properties.
func Accent(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"badge":        true,
				"badge-accent": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Ghost generates a ghost badge element based on the provided properties.
func Ghost(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"badge":       true,
				"badge-ghost": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Warning generates a warning badge element based on the provided properties.
func Warning(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"badge":         true,
				"badge-warning": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Info generates an info badge element based on the provided properties.
func Info(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"badge":      true,
				"badge-info": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Success generates a success badge element based on the provided properties.
func Success(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"badge":         true,
				"badge-success": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Error generates an error badge element based on the provided properties.
func Error(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"badge":       true,
				"badge-error": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
