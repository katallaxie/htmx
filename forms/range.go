package forms

import htmx "github.com/katallaxie/htmx"

// RangeProps represents the properties for a range input element.
type RangeProps struct {
	// Name is the name of the range input element.
	Name string
	// Value is the value of the range input element.
	Value string
	// Min is the minimum value of the range input element.
	Min string
	// Max is the maximum value of the range input element.
	Max string
	// Step is the step value of the range input element.
	Step string
	// Disabled indicates whether the range input is disabled.
	Disabled bool
	// Validator indicates whether the range input has validation styles.
	Validator bool
	// ClassNames are the additional class names for the range input element.
	htmx.ClassNames
}

// Range returns a range input element based on the provided properties.
func Range(p RangeProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"range":     true,
				"validator": p.Validator,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "range"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.Attribute("min", p.Min),
		htmx.Attribute("max", p.Max),
		htmx.Attribute("step", p.Step),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// RangeSuccess is a component that displays a success range input.
func RangeSuccess(p RangeProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"range":         true,
				"range-success": true,
				"validator":     p.Validator,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "range"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.Attribute("min", p.Min),
		htmx.Attribute("max", p.Max),
		htmx.Attribute("step", p.Step),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// RangeWarning is a component that displays a warning range input.
func RangeWarning(p RangeProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"range":         true,
				"range-warning": true,
				"validator":     p.Validator,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "range"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.Attribute("min", p.Min),
		htmx.Attribute("max", p.Max),
		htmx.Attribute("step", p.Step),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// RangeInfo is a component that displays an info range input.
func RangeInfo(p RangeProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"range":      true,
				"range-info": true,
				"validator":  p.Validator,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "range"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.Attribute("min", p.Min),
		htmx.Attribute("max", p.Max),
		htmx.Attribute("step", p.Step),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}

// RangeError is a component that displays an error range input.
func RangeError(p RangeProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"range":       true,
				"range-error": true,
				"validator":   p.Validator,
			},
			p.ClassNames,
		),
		htmx.Attribute("type", "range"),
		htmx.Attribute("name", p.Name),
		htmx.Attribute("value", p.Value),
		htmx.Attribute("min", p.Min),
		htmx.Attribute("max", p.Max),
		htmx.Attribute("step", p.Step),
		htmx.If(p.Disabled, htmx.Disabled()),
		htmx.Group(children...),
	)
}
