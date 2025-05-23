// Code generated by herogen; DO NOT EDIT.
package heroicons

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/icons"
)

func FaceSmilePlusOutline(p icons.IconProps) htmx.Node {
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
			htmx.Attribute("d", "M 18.5,3.0000001 V 8 M 21,5.5 h -5 m 4.863061,4.937375 A 9,9 0 0 1 16.155529,19.983306 9,9 0 0 1 5.6358301,18.36417 9,9 0 0 1 4.0166936,7.8444713 9,9 0 0 1 13.562625,3.1369391 M 15.182,15.182 a 4.5,4.5 0 0 1 -6.364,0 M 9.75,9.75 C 9.75,10.164 9.582,10.5 9.375,10.5 9.168,10.5 9,10.164 9,9.75 9,9.336 9.168,9 9.375,9 9.582,9 9.75,9.336 9.75,9.75 Z m -0.375,0 H 9.383 V 9.765 H 9.375 Z M 15,9.75 C 15,10.164 14.832,10.5 14.625,10.5 14.418,10.5 14.25,10.164 14.25,9.75 14.25,9.336 14.418,9 14.625,9 14.832,9 15,9.336 15,9.75 Z m -0.375,0 h 0.008 v 0.015 h -0.008 z"),
		),
	)
}
