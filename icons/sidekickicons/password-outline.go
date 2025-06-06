// Code generated by herogen; DO NOT EDIT.
package heroicons

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/icons"
)

func PasswordOutline(p icons.IconProps) htmx.Node {
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
			htmx.Attribute("d", "m 16,12.75 2.598076,-1.5 m 0,1.5 L 16,11.25 m 1.299038,-0.75 v 3 M 5.450962,12.75 8.049038,11.25 m 0,1.5 -2.598076,-1.5 M 6.75,10.5 v 3 m 3.950962,-0.75 2.598076,-1.5 m 0,1.5 -2.598076,-1.5 M 12,10.5 v 3 M 4.75,7.75 C 3.5073611,7.75 2.5,8.757361 2.5,10 v 4 c 0,1.242599 1.0073611,2.25 2.25,2.25 h 14.5 c 1.242599,0 2.25,-1.007401 2.25,-2.25 v -4 c 0,-1.242639 -1.007361,-2.25 -2.25,-2.25 z"),
		),
	)
}
