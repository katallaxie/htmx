package accordions

import htmx "github.com/katallaxie/htmx"

// Props is a component that can be expanded and collapsed.
type Props struct {
	ClassNames htmx.ClassNames
	Name       string
	Checked    bool
}

// Accordion is a component that can be expanded and collapsed.
func Accordion(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"bg-base-200":     true,
				"border-base-300": true,
				"border":          true,
				"collapse":        true,
			},
			props.ClassNames,
		),
		Radio(
			RadioProps{
				Name:    props.Name,
				Checked: props.Checked,
			},
		),
		htmx.Group(children...),
	)
}

// Arrow is a component that can be expanded and collapsed.
func Arrow(props Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"collapse":       true,
				"collapse-arrow": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// TitleProps is a component title.
type TitleProps struct {
	htmx.ClassNames
}

// Title is a component that can be expanded and collapsed.
func Title(props TitleProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"collapse-title": true,
				"font-semibold":  true,
			},
			props.ClassNames,
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

// RadioProps is a component that can be expanded and collapsed.
type RadioProps struct {
	// Name is the name of the radio button.
	Name string
	// Checked is the checked state of the radio button.
	Checked bool

	htmx.ClassNames
}

// Radio is a component that can be expanded and collapsed.
func Radio(props RadioProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Type("radio"),
		htmx.Name(props.Name),
		htmx.If(props.Checked, htmx.Checked()),
		htmx.Group(children...),
	)
}
