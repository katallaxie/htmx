package forms

import htmx "github.com/katallaxie/htmx"

// FieldsetProps contains the properties for the fieldset component.
type FieldsetProps struct {
	ClassNames htmx.ClassNames // The class names for the fieldset element.
}

// Fieldset is a component for the htmx fieldset extension.
func Fieldset(p FieldsetProps, children ...htmx.Node) htmx.Node {
	return htmx.Fieldset(
		htmx.Merge(
			htmx.ClassNames{
				"fieldset": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// FieldSetLegendProps contains the properties for the legend component.
type FieldSetLegendProps struct {
	ClassNames htmx.ClassNames // The class names for the legend element.
}

// FieldSetLegend is a component for the htmx legend extension.
func FieldSetLegend(p FieldSetLegendProps, children ...htmx.Node) htmx.Node {
	return htmx.Legend(
		htmx.Merge(
			htmx.ClassNames{
				"fieldset-legend": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
