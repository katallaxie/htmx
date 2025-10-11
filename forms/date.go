package forms

import (
	htmx "github.com/katallaxie/htmx"
)

// DateInputProps represents the properties for a date input element.
type DateInputProps struct {
	// Validator indicates whether the input should have validation.
	Validator bool

	htmx.ClassNames
}

// DateInput returns a date input element based on the provided properties.
func DateInput(p DateInputProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"validator": p.Validator,
			},
			p.ClassNames,
		),
		htmx.Type("date"),
		htmx.Group(children...),
	)
}
