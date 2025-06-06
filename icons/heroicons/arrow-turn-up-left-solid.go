// Code generated by herogen; DO NOT EDIT.
package heroicons

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/icons"
)

func ArrowTurnUpLeftSolid(p icons.IconProps) htmx.Node {
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
			htmx.Attribute("d", "M20.2388 20.25C19.8245 20.25 19.4887 19.9142 19.4887 19.5L19.4887 8.99902L5.54982 8.99902L8.0195 11.469C8.31238 11.7619 8.31236 12.2368 8.01944 12.5297C7.72653 12.8226 7.25165 12.8226 6.95877 12.5297L3.20889 8.77932C2.91603 8.48642 2.91603 8.01156 3.20889 7.71866L6.95877 3.9683C7.25165 3.67538 7.72653 3.67535 8.01944 3.96825C8.31236 4.26114 8.31238 4.73603 8.0195 5.02896L5.54982 7.49895L20.2388 7.49895C20.653 7.49895 20.9888 7.83476 20.9888 8.24899L20.9888 19.5C20.9888 19.9142 20.653 20.25 20.2388 20.25Z"),
		),
	)
}
