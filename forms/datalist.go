package forms

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/pkg/utilx"
)

// DatalistProps is the type of the props for the Datalist component.
type DatalistProps struct {
	// ID is the id of the datalist
	ID string
	// Target is the target of the datalist
	Target string
	// URL is the url of the datalist
	URL string
	// ClassNames is a map of class names to conditionally add to the component
	ClassNames htmx.ClassNames
	// Name is the name of the datalist
	Name string
	// Placeholder is the placeholder of the datalist
	Placeholder string
	// Disabled is whether the datalist is disabled
	Disabled bool
	// Indicator is the indicator of the datalist
	Indicator string
	// Error is the error of the datalist
	Error error
}

// Datalist is a component that displays a datalist.
func Datalist(props DatalistProps, children ...htmx.Node) htmx.Node {
	return htmx.Fragment(
		htmx.Script(htmx.Raw(`function checkUserKeydown(event) {return event instanceof KeyboardEvent}`)),
		TextInputBordered(
			TextInputProps{
				ClassNames:  props.ClassNames,
				Disabled:    props.Disabled,
				Error:       props.Error,
				Name:        props.Name,
				Placeholder: props.Placeholder,
			},
			htmx.List(props.ID),
			htmx.HxGet(props.URL),
			htmx.HxTarget(props.Target),
			htmx.HxSwap("innerHTML"),
			htmx.HxTrigger("load, keyup[checkUserKeydown.call(this, event)] changed delay:350ms"),
			htmx.HxIndicator(
				utilx.IfElse(
					utilx.Empty(props.Indicator),
					htmx.HxClssNameIndicatorSelector,
					props.Indicator,
				),
			),
			htmx.Group(children...),
		),
		htmx.DataList(
			htmx.ID(props.ID),
		),
	)
}
