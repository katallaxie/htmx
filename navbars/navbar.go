package navbars

import htmx "github.com/katallaxie/htmx"

// Props are the properties of the Navbar component.
type Props struct {
	htmx.ClassNames
}

// Navbar represents a Daisy UI Navbar component.
// see: https://daisyui.com/components/navbar
func Navbar(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"navbar": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// StartProps represents the properties of the NavbarStart component.
type StartProps struct {
	ClassNames htmx.ClassNames
}

// Start represents a Daisy UI NavbarStart component.
func Start(props StartProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"navbar-start": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// CenterProps are the properties of the NavbarCenter component.
type CenterProps struct {
	ClassNames htmx.ClassNames
}

// Center represents a Daisy UI NavbarCenter component.
func Center(props CenterProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"navbar-center": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// EndProps are the properties of the NavbarEnd component.
type EndProps struct {
	ClassNames htmx.ClassNames
}

// End represents a Daisy UI NavbarEnd component.
func End(props EndProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"navbar-end": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
