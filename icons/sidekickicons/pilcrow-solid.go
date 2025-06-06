// Code generated by herogen; DO NOT EDIT.
package heroicons

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/icons"
)

func PilcrowSolid(p icons.IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Merge(
			htmx.ClassNames{
				"h-5": false,
				"h-6": true,
				"w-5": false,
				"w-6": true,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("d", "M 10.5,3 C 8.0236046,3 6,5.0236046 6,7.5 6,9.9763954 8.0236046,12 10.5,12 h 1 v 8.25 c 0,0.414214 0.335786,0.75 0.75,0.75 0.414214,0 0.75,-0.335786 0.75,-0.75 v -9 -6.75 h 1.5 V 20.25 C 14.5,20.664214 14.835786,21 15.25,21 15.664214,21 16,20.664214 16,20.25 V 4.5 h 1.25 C 17.664214,4.5 18,4.1642136 18,3.75 18,3.3357864 17.664214,3 17.25,3 h -2 -3 z"),
		),
	)
}
