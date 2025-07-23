package navbars_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/navbars"
	"github.com/stretchr/testify/require"
)

func TestNavbar(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    navbars.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<div class=\"navbar\"></div>",
			props:    navbars.Props{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<div class=\"custom-class navbar\"></div>",
			props: navbars.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := navbars.Navbar(
				test.props,
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}

func TestStart(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    navbars.StartProps
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<div class=\"navbar-start\"></div>",
			props:    navbars.StartProps{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<div class=\"custom-class navbar-start\"></div>",
			props: navbars.StartProps{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := navbars.Start(
				test.props,
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}

func TestEnd(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    navbars.EndProps
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<div class=\"navbar-end\"></div>",
			props:    navbars.EndProps{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<div class=\"custom-class navbar-end\"></div>",
			props: navbars.EndProps{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := navbars.End(
				test.props,
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}

func TestCenter(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    navbars.CenterProps
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<div class=\"navbar-center\"></div>",
			props:    navbars.CenterProps{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<div class=\"custom-class navbar-center\"></div>",
			props: navbars.CenterProps{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := navbars.Center(
				test.props,
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}
