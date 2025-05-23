// Code generated by herogen; DO NOT EDIT.
package heroicons

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/icons"
)

func RobotOutline(p icons.IconProps) htmx.Node {
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
			htmx.Attribute("d", "m 8,17.750003 h 8 M 19.5,12.25 l 0.75,10e-7 a 1.5000022,1.5000022 0 0 1 1.5,1.500002 v 2 a 1.4999982,1.4999982 0 0 1 -1.5,1.499998 L 19.5,17.25 m -15,-5 -0.750003,10e-7 A 1.499999,1.499999 0 0 0 2.25,13.75 v 2.000003 a 1.4999982,1.4999982 0 0 0 1.5000002,1.499998 L 4.5,17.25 M 12,8.000003 v -2 m 1.5,-1.5000011 a 1.5,1.5 0 0 1 -1.5,1.5 1.5,1.5 0 0 1 -1.5,-1.5 1.5,1.5 0 0 1 1.5,-1.5 1.5,1.5 0 0 1 1.5,1.5 z m 2.75,8.7500011 a 1.5,1.5 0 0 1 -1.5,1.5 1.5,1.5 0 0 1 -1.5,-1.5 1.5,1.5 0 0 1 1.5,-1.5 1.5,1.5 0 0 1 1.5,1.5 z m -5.5,0 a 1.5,1.5 0 0 1 -1.5,1.5 1.5,1.5 0 0 1 -1.5,-1.5 1.5,1.5 0 0 1 1.5,-1.5 1.5,1.5 0 0 1 1.5,1.5 z M 8.2499952,8.5000029 h 7.5000098 c 2.0775,0 3.75,1.6725001 3.75,3.7500001 v 5 c 0,2.0775 -1.6725,3.75 -3.75,3.75 H 8.2499952 c -2.0775,0 -3.75,-1.6725 -3.75,-3.75 v -5 c 0,-2.0775 1.6725,-3.7500001 3.75,-3.7500001 z"),
		),
	)
}
