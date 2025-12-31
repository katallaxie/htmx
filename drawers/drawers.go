package drawers

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/pkg/utilx"
)

// Props is the props for the Drawer component.
type Props struct {
	// Open indicates whether the drawer is open.
	Open bool
	// ClassNames is a set of class names to apply to the drawer.
	htmx.ClassNames
}

// Drawer is a component that renders a drawer.
func Drawer(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"drawer":      true,
				"drawer-open": p.Open,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ToggleButtonProps is the props for the DrawerToggleButton component.
type ToggleButtonProps struct {
	// ID is the ID of the drawer toggle button.
	ID string
	// ClassNames is a set of class names to apply to the button.
	htmx.ClassNames
}

// ToggleButton is a component that renders a drawer toggle button.
func ToggleButton(p ToggleButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Label(
		htmx.Merge(
			htmx.ClassNames{
				"btn":           true,
				"drawer-button": true,
			},
			p.ClassNames,
		),
		htmx.For(p.ID),
		htmx.Group(children...),
	)
}

// ToggleProps is the props for the DrawerToggle component.
type ToggleProps struct {
	// ID is the ID of the drawer toggle button.
	ID string
	// AriaLabel is the aria-label of the drawer toggle button.
	AriaLabel string
	// ClassNames is a set of class names to apply to the button.
	htmx.ClassNames
}

// Toggle is a component that renders a drawer toggle button.
func Toggle(p ToggleProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"drawer-toggle": true,
			},
			p.ClassNames,
		),
		htmx.ID(p.ID),
		htmx.Type("checkbox"),
		htmx.If(utilx.NotEmpty(p.AriaLabel), htmx.AriaLabel(p.AriaLabel)),
		htmx.Group(children...),
	)
}

// ContentProps is the props for the DrawerContent component.
type ContentProps struct {
	// ClassNames is a set of class names to apply to the drawer content.
	htmx.ClassNames
}

// Content is a component that renders the content of a drawer.
func Content(p ContentProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"drawer-content": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// SideProps is the props for the DrawerSide component.
type SideProps struct {
	// ClassNames is a set of class names to apply to the drawer side.
	htmx.ClassNames
}

// Side is a component that renders the side of a drawer.
func Side(p SideProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"drawer-side": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
