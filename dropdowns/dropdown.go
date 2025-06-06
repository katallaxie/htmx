package dropdowns

import (
	htmx "github.com/katallaxie/htmx"

	"github.com/katallaxie/pkg/conv"
)

// DropdownProps represents the properties for a dropdown element.
type DropdownProps struct {
	htmx.ClassNames // The class names for the dropdown element.
}

// Dropdown generates a dropdown element based on the provided properties.
func Dropdown(p DropdownProps, children ...htmx.Node) htmx.Node {
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

// DropdownButtonProps represents the properties for a dropdown summary element.
type DropdownButtonProps struct {
	TabIndex int
	htmx.ClassNames
}

// DropdownButton generates a dropdown summary element based on the provided properties.
func DropdownButton(p DropdownButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Summary(
		htmx.TabIndex(conv.String(p.TabIndex)),
		htmx.Role("button"),
		htmx.Merge(
			htmx.ClassNames{
				"btn": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// DropdownContentProps represents the properties for a dropdown content element.
type DropdownContentProps struct {
	htmx.ClassNames // The class names for the dropdown content element.
}

// DropdownContent generates a dropdown content element based on the provided properties.
func DropdownContent(p DropdownContentProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"dropdown-content": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// DropdownMenuItemsProps represents the properties for a dropdown menu items element.
type DropdownMenuItemsProps struct {
	TabIndex int
	htmx.ClassNames
}

// DropdownMenuItems generates a dropdown menu items element based on the provided properties.
func DropdownMenuItems(p DropdownMenuItemsProps, children ...htmx.Node) htmx.Node {
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
				"z-[1]":            true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// DropdownMenuItem represents the properties for a dropdown items element.
type DropdownMenuItemProps struct {
	htmx.ClassNames // The class names for the dropdown items element.
}

// DropdownMenuItem generates a dropdown items element based on the provided properties.
func DropdownMenuItem(p DropdownMenuItemProps, children ...htmx.Node) htmx.Node {
	return htmx.Li(
		htmx.Merge(
			htmx.ClassNames{},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
