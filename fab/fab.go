package fab

import (
	"github.com/katallaxie/htmx"
	"github.com/katallaxie/pkg/conv"
	"github.com/katallaxie/pkg/utilx"
)

// Props represents the properties for a floating action button (FAB).
type Props struct {
	htmx.ClassNames
}

// Fab generates a floating action button (FAB) based on the provided properties.
func Fab(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"fab": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ButtonProps represents the properties for a button inside the floating action button (FAB).
type ButtonProps struct {
	// TabIndex sets the tabindex attribute.
	TabIndex int
	// Role sets the role attribute.
	Role string

	htmx.ClassNames
}

// Button generates a button for the floating action button (FAB) based on the provided properties.
func Button(p ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"btn": true,
			},
			p.ClassNames,
		),
		htmx.IfElse(utilx.NotEmpty(p.Role), htmx.Role(p.Role), htmx.Role("button")),
		htmx.TabIndex(conv.String(p.TabIndex)),
		htmx.Group(children...),
	)
}
