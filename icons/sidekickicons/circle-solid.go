// Code generated by herogen; DO NOT EDIT.
package heroicons

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/icons"
)

func CircleSolid(p icons.IconProps) htmx.Node {
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
			htmx.Attribute("d", "m 12,2.25 c -5.38478,0 -9.75,4.36522 -9.75,9.75 0,5.3848 4.36522,9.75 9.75,9.75 5.3848,0 9.75,-4.3652 9.75,-9.75 0,-5.38478 -4.3652,-9.75 -9.75,-9.75 z"),
		),
	)
}
