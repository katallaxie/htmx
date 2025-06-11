package modals

import htmx "github.com/katallaxie/htmx"

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

// ModalBox contains the properties for the modal box component.
type ModalBoxProps struct {
	htmx.ClassNames
}

// ModalBox is a component for the htmx modal extension.
func ModalBox(p ModalBoxProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"modal-box": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ModalActionProps contains the properties for the modal actions component.
type ModalActionProps struct {
	// ClassNames are the CSS class names to apply to the modal actions.
	htmx.ClassNames
}

// ModalAction is a component for the htmx modal extension.
func ModalAction(p ModalActionProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"modal-action": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ModalCloseButtonProps contains the properties for the modal close button component.
type ModalCloseButtonProps struct {
	// ClassNames are the CSS class names to apply to the modal close button.
	htmx.ClassNames
}

// ModalCloseButton is a component for the htmx modal extension.
func ModalCloseButton(p ModalCloseButtonProps, children ...htmx.Node) htmx.Node {
	return htmx.Form(
		htmx.Method("dialog"),
		htmx.Button(
			htmx.Merge(
				htmx.ClassNames{
					"btn": true,
				},
				p.ClassNames,
			),
			htmx.Text("Close"),
		),
		htmx.Group(children...),
	)
}
