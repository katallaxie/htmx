// Code generated by herogen; DO NOT EDIT.
package heroicons

import (
	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/icons"
)

func BadgeSolid(p icons.IconProps) htmx.Node {
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
			htmx.Attribute("d", "M 8.60288,3.79876 C 9.42705,2.85093 10.6433,2.25 12,2.25 c 1.3566,0 2.5728,0.60087 3.397,1.54861 1.2531,-0.08755 2.5382,0.34755 3.4976,1.30696 0.9594,0.95941 1.3945,2.24452 1.307,3.49762 C 21.1492,9.42735 21.75,10.6435 21.75,12 c 0,1.3568 -0.601,2.5731 -1.5489,3.3973 0.0873,1.2529 -0.3478,2.5377 -1.307,3.497 -0.9593,0.9592 -2.2441,1.3943 -3.497,1.307 C 14.5729,21.1491 13.3567,21.75 12,21.75 10.6434,21.75 9.4272,21.1491 8.60304,20.2014 7.34992,20.289 6.0648,19.8539 5.10537,18.8945 4.14596,17.935 3.71086,16.6499 3.79841,15.3968 2.85079,14.5726 2.25,13.3565 2.25,12 2.25,10.6434 2.85085,9.42723 3.79855,8.60306 3.7111,7.35005 4.14621,6.06507 5.10554,5.10574 6.06488,4.1464 7.34987,3.71129 8.60288,3.79876 Z"),
		),
	)
}
