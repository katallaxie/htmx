// Code generated by herogen; DO NOT EDIT.
package heroicons

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/icons"
)

func QuotationMarkOutline(p icons.IconProps) htmx.Node {
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
			htmx.Attribute("d", "M 16.250073,15.5 C 14.412507,14.705819 13.75,14.000001 13.75,12.000001 c 0,-2 1.250146,-3.250002 3.250073,-3.250001 C 19,8.750001 20.25,10.000001 20.25,12.000001 c 0,3.5 -0.999927,4.749999 -4.999927,7.249999 1,-1.75 1,-2.25 1,-3.75 z m -10,0 C 4.412507,14.705819 3.75,14.000001 3.75,12.000001 3.75,10.000001 5.000146,8.749999 7.000073,8.75 9,8.750001 10.25,10.000001 10.25,12.000001 c 0,3.5 -0.999927,4.749999 -4.999927,7.249999 1,-1.75 1,-2.25 1,-3.75 z"),
		),
	)
}
