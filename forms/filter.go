package forms

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/pkg/utilx"
)

// FilterProps is a struct that contains the properties of the Filter component.
type FilterProps struct {
	// ClassNames are the additional class names for the filter.
	htmx.ClassNames
}

// Filter is a component that renders a filter element.
func Filter(props FilterProps, children ...htmx.Node) htmx.Node {
	return htmx.Form(
		htmx.Merge(
			htmx.ClassNames{
				"filter": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// FilterResetButtonProps is a struct that contains the properties of the FilterReset component.
type FilterResetProps struct {
	// Value is the value of the reset button.
	Value string
	// ClassNames are the additional class names for the reset button.
	htmx.ClassNames
}

// FilterReset is a component that renders a filter reset button.
func FilterReset(props FilterResetProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"btn":        true,
				"btn-square": true,
			},
			props.ClassNames,
		),
		htmx.Type("reset"),
		utilx.IfElse(utilx.NotEmpty(props.Value), htmx.Value(props.Value), htmx.Value("×")),
		htmx.Group(children...),
	)
}

// FilterOptionProps is a struct that contains the properties of the FilterOption component.
type FilterOptionProps struct {
	// Name is the name of the filter option.
	Name string
	// AriaLabel is the aria-label of the filter option.
	AriaLabel string
	// ClassNames are the additional class names for the filter option.
	htmx.ClassNames
}

// FilterOption is a component that renders a filter option.
func FilterOption(props FilterOptionProps, children ...htmx.Node) htmx.Node {
	return htmx.Input(
		htmx.Merge(
			htmx.ClassNames{
				"btn": true,
			},
			props.ClassNames,
		),
		htmx.Type("checkbox"),
		htmx.Name(props.Name),
		htmx.AriaLabel(props.AriaLabel),
		htmx.Group(children...),
	)
}
