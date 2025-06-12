package drawers

import htmx "github.com/katallaxie/htmx"

// Props is the props for the Drawer component.
type Props struct {
	// ClassNames is a set of class names to apply to the drawer.
	htmx.ClassNames
}

// Drawer is a component that renders a drawer.
func Drawer(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"drawer": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ToggleProps is the props for the DrawerToggle component.
type ToggleProps struct {
	// ID is the ID of the drawer toggle button.
	ID string
	// ClassNames is a set of class names to apply to the button.
	htmx.ClassNames
}

// DrawerToggle is a component that renders a drawer toggle button.
func DrawerToggle(p ToggleProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"drawer-toggle": true,
			},
			p.ClassNames,
		),
		htmx.ID(p.ID),
		htmx.Type("checkbox"),
		htmx.Group(children...),
	)
}

// DrawerContentProps is the props for the DrawerContent component.
type DrawerContentProps struct {
	// ClassNames is a set of class names to apply to the drawer content.
	htmx.ClassNames
}

// DrawerContent is a component that renders the content of a drawer.
func DrawerContent(p DrawerContentProps, children ...htmx.Node) htmx.Node {
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

// DrawerSideProps is the props for the DrawerSide component.
type DrawerSideProps struct {
	// ClassNames is a set of class names to apply to the drawer side.
	htmx.ClassNames
}

// DrawerSide is a component that renders the side of a drawer.
func DrawerSide(p DrawerSideProps, children ...htmx.Node) htmx.Node {
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

// DrawerOpenProps is the props for the DrawerOpen component.
type DrawerOpenProps struct {
	// ID is the ID of the drawer open button.
	ID string
	// ClassNames is a set of class names to apply to the button.
	htmx.ClassNames
}

// DrawerOpen is a component that renders a drawer open button.
func DrawerOpen(p DrawerOpenProps, children ...htmx.Node) htmx.Node {
	return htmx.Label(
		htmx.Merge(
			htmx.ClassNames{
				"btn":           true,
				"btn-primary":   true,
				"drawer-button": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("for", p.ID),
		htmx.Group(children...),
	)
}

// DrawerCloseProps is the props for the DrawerClose component.
type DrawerCloseProps struct {
	// ID is the ID of the drawer close button.
	ID string
	// ClassNames is a set of class names to apply to the button.
	htmx.ClassNames
}

// DrawerClose is a component that renders a drawer close button.
func DrawerClose(p DrawerCloseProps, children ...htmx.Node) htmx.Node {
	return htmx.Label(
		htmx.Merge(
			htmx.ClassNames{
				"btn":          true,
				"btn-ghost":    true,
				"drawer-close": true,
			},
			p.ClassNames,
		),
		htmx.Attribute("for", p.ID),
		htmx.Group(children...),
	)
}
