package collapsible

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/pkg/conv"
)

// Props is a component that can be expanded and collapsed.
type Props struct {
	// TabIndex sets the tabindex attribute.
	TabIndex int

	htmx.ClassNames
}

// Collapse is a component that can be expanded and collapsed.
func Collapse(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"collapse": true,
			},
			props.ClassNames,
		),
		htmx.Attribute("tabindex", conv.String(props.TabIndex)),
		htmx.Group(children...),
	)
}

// Arrow is a component that can be expanded and collapsed.
func Arrow(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Attribute("tabindex", conv.String(p.TabIndex)),
		htmx.Merge(
			htmx.ClassNames{
				"collapse":        true,
				"collapse-arrow":  true,
				"bg-base-200":     true,
				"border":          true,
				"border-base-300": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// TitleProps is a component that can be expanded and collapsed.
type TitleProps struct {
	htmx.ClassNames
}

// Title is a component that can be expanded and collapsed.
func Title(p TitleProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"collapse-title": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ContentProps is a component that can be expanded and collapsed.
type ContentProps struct {
	htmx.ClassNames
}

// Content is a component that can be expanded and collapsed.
func Content(props ContentProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"collapse-content": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// CheckboxProps is a component that can be expanded and collapsed.
type CheckboxProps struct {
	Checked bool

	htmx.ClassNames
}

// Checkbox is a component that can be expanded and collapsed.
func Checkbox(props CheckboxProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Group(children...),
		htmx.Merge(props.ClassNames),
		htmx.Type("checkbox"),
	)
}
