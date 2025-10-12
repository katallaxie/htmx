package avatars

import htmx "github.com/katallaxie/htmx"

// Props represents the properties for an avatar.
type Props struct {
	htmx.ClassNames
}

// Avatar generates an avatar based on the provided properties.
func Avatar(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Group generates a group of avatars based on the provided properties.
func Group(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar-group": true,
				"-space-x-6":   true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Online generates an online avatar based on the provided properties.
func Online(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar":        true,
				"avatar-online": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Offline generates an offline avatar based on the provided properties.
func Offline(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar":         true,
				"avatar-offline": true,
			},
			p.ClassNames,
		),
		htmx.Group(children...),
	)
}

// Placeholder generates a placeholder for an avatar based on the provided properties.
func Placeholder(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.ClassNames{
			"avatar":             true,
			"avatar-placeholder": true,
		},
		htmx.Div(
			htmx.Merge(
				htmx.ClassNames{
					"bg-neutral":           true,
					"w-12":                 true,
					"text-neutral-content": true,
				},
				p.ClassNames,
			),
			htmx.Group(children...),
		),
	)
}

// RoundedMedium generates an avatar based on the provided properties.
func RoundedMedium(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar": true,
			},
			p.ClassNames,
		),
		htmx.Div(
			htmx.ClassNames{
				"w-24":    true,
				"rounded": true,
			},
			htmx.Group(children...),
		),
	)
}

// RoundedSmall generates an extra small avatar based on the provided properties.
func RoundedSmall(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar": true,
			},
			p.ClassNames,
		),
		htmx.Div(
			htmx.ClassNames{
				"w-8":     true,
				"rounded": true,
			},
			htmx.Group(children...),
		),
	)
}

// RoundedLarge generates a large avatar based on the provided properties.
func RoundedLarge(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar": true,
			},
			p.ClassNames,
		),
		htmx.Div(
			htmx.ClassNames{
				"w-32":    true,
				"rounded": true,
			},
			htmx.Group(children...),
		),
	)
}

// RoundSmall generates an extra small round avatar based on the provided properties.
func RoundSmall(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar": true,
			},
			p.ClassNames,
		),
		htmx.Div(
			htmx.ClassNames{
				"w-8":          true,
				"rounded-full": true,
			},
			htmx.Group(children...),
		),
	)
}

// RoundLarge generates a large round avatar based on the provided properties.
func RoundLarge(p Props, children ...htmx.Node) htmx.Node {
	return htmx.Div(
		htmx.Merge(
			htmx.ClassNames{
				"avatar": true,
			},
			p.ClassNames,
		),
		htmx.Div(
			htmx.ClassNames{
				"w-32":         true,
				"rounded-full": true,
			},
			htmx.Group(children...),
		),
	)
}
