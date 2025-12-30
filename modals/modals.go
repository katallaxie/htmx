package modals

import (
	htmx "github.com/katallaxie/htmx"
)

// Props contains the properties for the modal component.
type Props struct {
	// ID is the identifier for the modal.
	ID string
	// ClassNames are the CSS class names to apply to the modal.
	htmx.ClassNames
}

// Modal is a component for the htmx modal extension.
func Modal(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Dialog(
		htmx.Merge(
			htmx.ClassNames{
				"modal": true,
			},
			p.ClassNames,
		),
		htmx.ID(p.ID),
		htmx.Group(children...),
	)
}

// Box contains the properties for the modal box component.
type BoxProps struct {
	htmx.ClassNames
}

// Box is a component for the htmx modal extension.
func Box(props BoxProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"modal-box": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ActionProps contains the properties for the modal actions component.
type ActionProps struct {
	// ClassNames are the CSS class names to apply to the modal actions.
	htmx.ClassNames
}

// Action is a component for the htmx modal extension.
func Action(props ActionProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"modal-action": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}

// CloseButtonProps contains the properties for the modal close button component.
type CloseButtonProps struct {
	// ClassNames are the CSS class names to apply to the modal close button.
	htmx.ClassNames
	// ID is the identifier for the modal close button.
	ID string
}

// CloseButton is a component for the htmx modal extension.
func CloseButton(props CloseButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn": true,
			},
			props.ClassNames,
		),
		htmx.CommandFor(props.ID),
		htmx.Command("close"),
		htmx.Group(children...),
	)
}

// OpenButtonProps contains the properties for the modal open button component.
type OpenButtonProps struct {
	// ClassNames are the CSS class names to apply to the modal open button.
	htmx.ClassNames
	// ID is the identifier for the modal open button.
	ID string
}

// OpenButton is a component for the htmx modal extension.
func OpenButton(props OpenButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Button(
		htmx.Merge(
			htmx.ClassNames{
				"btn": true,
			},
			props.ClassNames,
		),
		htmx.CommandFor(props.ID),
		htmx.Command("show-modal"),
		htmx.Group(children...),
	)
}
