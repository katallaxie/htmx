package dropdowns

import (
	htmx "github.com/katallaxie/htmx"

	"github.com/katallaxie/pkg/conv"
)

// Props represents the properties for a dropdown element.
type Props struct {
	htmx.ClassNames // The class names for the dropdown element.
}

// Dropdown generates a dropdown element based on the provided properties.
func Dropdown(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Details(
		htmx.Merge(
			htmx.ClassNames{
				"dropdown": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ButtonProps represents the properties for a dropdown summary element.
type ButtonProps struct {
	TabIndex int
	htmx.ClassNames
}

// Button generates a dropdown summary element based on the provided properties.
func Button(p ButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Summary(
		htmx.If(p.TabIndex > 0, htmx.TabIndex(conv.String(p.TabIndex))),
		htmx.Merge(
			htmx.ClassNames{
				"btn": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// MenuItemsProps represents the properties for a dropdown menu items element.
type MenuItemsProps struct {
	TabIndex int
	htmx.ClassNames
}

// MenuItems generates a dropdown menu items element based on the provided properties.
func MenuItems(p MenuItemsProps, children ...htmx.Node) htmx.Node {
	return htmx.Ul(
		htmx.TabIndex(
			conv.String(p.TabIndex),
		),
		htmx.Merge(
			htmx.ClassNames{
				"bg-base-100":      true,
				"dropdown-content": true,
				"menu":             true,
				"p-2":              true,
				"rounded-box":      true,
				"shadow":           true,
				"w-52":             true,
				"z-1":              true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// MenuItem represents the properties for a dropdown items element.
type MenuItemProps struct {
	htmx.ClassNames // The class names for the dropdown items element.
}

// MenuItem generates a dropdown items element based on the provided properties.
func DropdownMenuItem(p MenuItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Li(
		htmx.Merge(
			htmx.ClassNames{},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
