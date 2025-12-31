package skeletons

import (
	"fmt"

	htmx "github.com/katallaxie/htmx"
)

// Props is a struct that contains the props of the skeleton component.
type Props struct {
	// Width describes the width of the skeleton.
	Width int
	// Height describes the height of the skeleton.
	Height int
	// ClassNames are the additional class names for the skeleton.
	htmx.ClassNames
}

// Skeleton is a component that renders a skeleton element.
func Skeleton(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"skeleton":                    true,
				fmt.Sprintf("w-%d", p.Width):  true,
				fmt.Sprintf("h-%d", p.Height): true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// SkeletonTextProps is a struct that contains the props of the skeleton text component.
type SkeletonTextProps struct {
	htmx.ClassNames
}

// SkeletonText is a component that renders a skeleton text element.
func SkeletonText(props SkeletonTextProps, children ...htmx.Node) htmx.Node {
	return htmx.Span(
		htmx.Merge(
			htmx.ClassNames{
				"skeleton":      true,
				"skeleton-text": true,
			},
			props.ClassNames,
		),
		htmx.Group(children...),
	)
}
