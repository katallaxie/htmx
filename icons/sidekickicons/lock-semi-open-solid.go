// Code generated by herogen; DO NOT EDIT.
package heroicons

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/icons"
)

func LockSemiOpenSolid(p icons.IconProps) htmx.Node {
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
			htmx.Attribute("d", "M 12 0.25 C 9.1093785 0.25 6.75 2.6093992 6.75 5.5 L 6.75 7 A 0.75 0.75 0 0 0 7.5 7.75 A 0.75 0.75 0 0 0 8.25 7 L 8.25 5.5 C 8.25 3.4200507 9.9200315 1.75 12 1.75 C 14.079969 1.75 15.75 3.4200507 15.75 5.5 L 15.75 9.75 L 6.75 9.75 C 5.1020166 9.75 3.75 11.10206 3.75 12.75 L 3.75 19.5 C 3.75 21.14794 5.1020166 22.5 6.75 22.5 L 17.25 22.5 C 18.897956 22.5 20.25 21.147956 20.25 19.5 L 20.25 12.75 C 20.25 11.102044 18.897956 9.75 17.25 9.75 L 17.25 5.5 C 17.25 2.6093992 14.890621 0.25 12 0.25 z "),
		),
	)
}
