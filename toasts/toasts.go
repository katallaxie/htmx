package toasts

import htmx "github.com/katallaxie/htmx"

// ToastProps are the properties for a toast component.
type ToastProps struct {
	htmx.ClassNames
}

// Toat is a function that creates a toast component.
func Toat(p ToastProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"toast": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ToastTop is a function that creates a toast component that is positioned at the top.
func ToastTop(p ToastProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"toast":     true,
				"toast-top": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ToastBottom is a function that creates a toast component that is positioned at the bottom.
func ToastBottom(p ToastProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"toast":        true,
				"toast-bottom": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ToastStart is a function that creates a toast component that is positioned at the start.
func ToastStart(p ToastProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"toast":       true,
				"toast-start": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ToastEnd is a function that creates a toast component that is positioned at the end.
func ToastEnd(p ToastProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"toast":     true,
				"toast-end": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ToastCenter is a function that creates a toast component that is positioned at the center.
func ToastCenter(p ToastProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"toast":        true,
				"toast-center": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// ToastMiddle is a function that creates a toast component that is positioned in the middle.
func ToastMiddle(p ToastProps, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"toast":        true,
				"toast-middle": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}
