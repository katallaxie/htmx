package drawers_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/drawers"
	"github.com/stretchr/testify/require"
)

func Test_Drawer(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    drawers.Props
	}{
		{
			name:     "default drawer",
			expected: `<div class="drawer"></div>`,
			children: nil,
			props:    drawers.Props{},
		},
		{
			name:     "drawer with class",
			expected: `<div class="custom-class drawer"></div>`,
			children: nil,
			props: drawers.Props{
				ClassNames: htmx.ClassNames{
					"custom-class": true,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := drawers.Drawer(
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

func Test_DrawerToggle(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    drawers.ToggleProps
	}{
		{
			name:     "default drawer toggle",
			expected: `<input class="drawer-toggle" id="" type="checkbox">`,
			children: nil,
			props:    drawers.ToggleProps{},
		},
		{
			name:     "drawer toggle with ID and class",
			expected: `<input class="custom-class drawer-toggle" id="toggle-id" type="checkbox">`,
			children: nil,
			props: drawers.ToggleProps{
				ID: "toggle-id",
				ClassNames: htmx.ClassNames{
					"custom-class": true,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := drawers.DrawerToggle(
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

func Test_DrawerContent(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    drawers.DrawerContentProps
	}{
		{
			name:     "default drawer content",
			expected: `<div class="drawer-content"></div>`,
			children: nil,
			props:    drawers.DrawerContentProps{},
		},
		{
			name:     "drawer content with class",
			expected: `<div class="custom-class drawer-content"></div>`,
			children: nil,
			props: drawers.DrawerContentProps{
				ClassNames: htmx.ClassNames{
					"custom-class": true,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := drawers.DrawerContent(
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
