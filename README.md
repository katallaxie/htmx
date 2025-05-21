# 🔨 HTMX

[![Test & Build](https://github.com/katallaxie/htmx/actions/workflows/main.yml/badge.svg)](https://github.com/katallaxie/htmx/actions/workflows/main.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/katallaxie/htmx.svg)](https://pkg.go.dev/github.com/katallaxie/htmx)
[![Go Report Card](https://goreportcard.com/badge/github.com/katallaxie/htmx)](https://goreportcard.com/report/github.com/katallaxie/htmx)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)

A package to write HTML5 and HTMX components in Go.

## 🦄 Features

- Write declartive HTML5 components in Go without using templates and with the full-power of a type-safe language, auto-completion, and refactoring.
- Full support for HTMX components.
- No dependencies on JavaScript frameworks.
- Fast rendering of HTML5 and HTMX components.
- Easy to use and learn.
- Easy to extend and customize.

## ✨ Components

There are additional complex components that help to write HTML5 and HTMX components in Go.

- [x] [HTMX](https://htmx.org/)
- [x] [HTML5](https://www.w3.org/TR/2011/WD-html5-20110405/)
- [x] [Heroicons](https://heroicons.com/)
- [x] [sidekickicons](https://sidekickicons.com)
- [x] [TailwindCSS](https://tailwindcss.com/)
- [ ] [DaisyUI](https://daisyui.com/) (83% 📈)

## 🛸 Installation

```bash
go get github.com/katallaxie/htmx
```

## 🧪 Usage

Creating a button leveraging htmx is as easy as this.

```go
htmx.Button(
    htmx.Attribute("type", "submit")
    htmx.Text("Button"),
    htmx.HxPost("/api/respond")
)
```

All nodes implement the `io.Writer` interface. Which means it is possible to render to a `http.Request`
but also to render to a file.

```go
f, err := os.OpenFile("123.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
if err != nil {
    panic(err)
}
defer f.Close()

err = Page().Render(f)
if err != nil {
    panic(err)
}
```

## 🎨 Elements

HTML and HTMX elements are represented as functions in Go. The functions are used to create the elements.

```go
htmx.Div(
    htmx.ClassNames{
        tailwind.FontSemibold: true,
    },
    htmx.Text("Hello World"),
)
```

This will create the following HTML element.

```html
<div class="font-semibold">Hello World</div>
```

There is support for all HTML5 elements and Tailwind classes. Use `import "github.com/katallaxie/fiber-htmx/tailwind"` to include Tailwind classes.

## 📦 Components

Write HTML5 and HTMX components in Go.

```go
func HelloWorld() htmx.Node {
    return htmx.Div(
        htmx.ClassNames{
            "font-semibold",
        },
        htmx.Text("Hello World"),
    )
}
```

There are different types of composition. For example, passing children to a component.

```go
func HelloWorld(children ...htmx.Node) htmx.Node {
    return htmx.Div(
        htmx.ClassNames{
            "font-semibold",
        },
        htmx.Text("Hello World"),
        htmx.Div(
            htmx.ClassNames{
                "text-red-500",
            },
            htmx.Group(children...),
        ),
    )
}
```

Styling of components is done with the `htmx.ClassNames` type.

```go
func HelloWorld() htmx.Node {
    return htmx.Div(
        htmx.ClassNames{
            tailwind.FontSemibold: true,
            "text-red-500": true,
        },
        htmx.Text("Hello World"),
    )
}
```

There are also helpers to make the life with styling easier by merging classes.

```go
func HelloWorld(classes htmx.ClassNames) htmx.Node {
    return htmx.Div(
        htmx.Merge(
            htmx.ClassNames{
                "font-semibold",
                "text-red-500",
            },
            classes,
        )
        htmx.Text("Hello World"),
    )
}
```

There is alos another pattern to create a component. This enables you to track state and to use the component in a more declarative way.

```go
type Page struct {
	htmx.Node
}

func NewPage(title, body string) *Page {
	return &Page{
		Node: htmx.HTML5(
			htmx.HTML5Props{
				Title: title,
			},
			htmx.Body(
				htmx.Div(
					htmx.Text(body),
				),
			),
		),
	}
}
```

## 🧩 Import map

An import map is a JSON object that allows developers to control how the browser resolves module specifiers when importing JavaScript modules.

```go
htmx.Imports(
    htmx.ImportsProp{
        Resolver: cache.New(jsdeliver.New()),
        Pkgs: []imports.ExactPackage{
            {
                Name:    "htmx.org",
                Version: "2.0.4",
            },
        },
        Requires: []imports.Require{
            {
                File: "dist/htmx.esm.js",
            },
        },
    }
),
```

[Import maps](https://github.com/WICG/import-maps) let you import JavaScript modules using logical names that map to versioned/digested files – directly from the browser. So you can [build modern JavaScript applications using JavaScript libraries made for ES modules (ESM) without the need for transpiling or bundling](https://world.hey.com/dhh/modern-web-apps-without-javascript-bundling-or-transpiling-a20f2755). This frees you from needing Webpack, Yarn, npm, or any other part of the JavaScript toolchain. All you need is the asset pipeline that's already included in Rails.

## 🏄‍♀️ Icons

The package has support for [Heroicons](https://heroicons.com/). The support is for the outline and solid icons.

```go
icon := heroicons.AcademicCapOutline(heroicons.IconProps{})
icon.Render(os.Stdout)
```

## 📄 Examples

See [examples](https://github.com/katallaxie/fiber-htmx/tree/master/examples) to understand the provided interfaces.

## 🏎️ Benchmarks

```bash
make bench
```

## License

[MIT](/LICENSE)
