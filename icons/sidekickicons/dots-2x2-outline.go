// Code generated by herogen; DO NOT EDIT.
package heroicons

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/icons"
)

func Dots2X2Outline(p icons.IconProps) htmx.Node {
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
			htmx.Attribute("d", "m 8.25,15 c 0,-0.4142 0.3358,-0.75 0.75,-0.75 0.4142,0 0.75,0.3358 0.75,0.75 0,0.4142 -0.3358,0.75 -0.75,0.75 -0.4142,0 -0.75,-0.3358 -0.75,-0.75 z m 6,0 c 0,-0.4142 0.3358,-0.75 0.75,-0.75 0.4142,0 0.75,0.3358 0.75,0.75 0,0.4142 -0.3358,0.75 -0.75,0.75 -0.4142,0 -0.75,-0.3358 -0.75,-0.75 z m -6,-6 C 8.25,8.5858 8.5858,8.25 9,8.25 9.4142,8.25 9.75,8.5858 9.75,9 9.75,9.4142 9.4142,9.75 9,9.75 8.5858,9.75 8.25,9.4142 8.25,9 Z m 6,0 c 0,-0.4142 0.3358,-0.75 0.75,-0.75 0.4142,0 0.75,0.3358 0.75,0.75 0,0.4142 -0.3358,0.75 -0.75,0.75 -0.4142,0 -0.75,-0.3358 -0.75,-0.75 z"),
		),
	)
}
