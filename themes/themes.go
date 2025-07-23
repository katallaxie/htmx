package themes

import htmx "github.com/katallaxie/htmx"

// Theme constants holds the names of various themes that can be applied to a web application.
const (
	Abyss        = "abyss"
	Acid         = "acid"
	Aqua         = "aqua"
	Autumn       = "autumn"
	Black        = "black"
	Bumblebee    = "bumblebee"
	Business     = "business"
	Caramellatte = "caramelatte"
	Coffee       = "coffee"
	Corporate    = "corporate"
	Cupcake      = "cupcake"
	CYMK         = "cmyk"
	Dark         = "dark"
	Dim          = "dim"
	Dracula      = "dracula"
	Emerald      = "emerald"
	Fantasy      = "fantasy"
	Forest       = "forest"
	Garden       = "garden"
	Halloween    = "halloween"
	Lemonade     = "lemonade"
	Light        = "light"
	Lofi         = "lofi"
	Luxury       = "luxury"
	Night        = "night"
	Nord         = "nord"
	Pastel       = "pastel"
	Retro        = "retro"
	Silk         = "silk"
	Sunset       = "sunset"
	Synthwave    = "synthwave"
	Valentine    = "valentine"
	Winter       = "winter"
	Wireframe    = "wireframe"
)

// Theme creates an htmx.Node that sets the data-theme attribute for theming purposes.
func Theme(theme string) htmx.Node {
	return htmx.Attribute("data-theme", theme)
}
