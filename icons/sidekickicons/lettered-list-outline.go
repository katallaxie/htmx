// Code generated by herogen; DO NOT EDIT.
package heroicons

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/icons"
)

func LetteredListOutline(p icons.IconProps) htmx.Node {
	return htmx.SVG(
		htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
		htmx.Attribute("viewBox", "0 0 24 24"),
		htmx.Merge(
			htmx.ClassNames{
				"h-5": true,
				"h-6": false,
				"w-5": true,
				"w-6": false,
			},
			p.ClassNames,
		),
		htmx.Path(
			htmx.Attribute("stroke-linecap", "round"),
			htmx.Attribute("stroke-linejoin", "round"),
			htmx.Attribute("d", "M 3,6 H 5.5 M 3,7.75 V 4.5 A 1.25,1.25 0 0 1 4.25,3.25 1.25,1.25 0 0 1 5.5,4.5 v 3.25 m 0,8.5 h -1 A 1.5,1.5 0 0 0 3,17.75 v 1.5 a 1.5,1.5 0 0 0 1.5,1.5 h 1 M 3,9.75 v 4.5 H 4.375 A 1.125,1.125 0 0 0 5.5,13.125 1.125,1.125 0 0 0 4.375,12 H 3.5 4.375 A 1.125,1.125 0 0 0 5.5,10.875 1.125,1.125 0 0 0 4.375,9.75 Z M 8.25,6 h 12 M 8.25046,12 H 20.25018 M 8.25,18 h 12"),
		),
	)
}
