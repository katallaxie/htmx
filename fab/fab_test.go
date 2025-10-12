package fab_test

import (
	"bytes"
	"testing"

	"github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/fab"
	"github.com/stretchr/testify/require"
)

func TestFabv(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    fab.Props
	}{
		{
			name:     "default",
			props:    fab.Props{},
			expected: `<div class="fab"></div>`,
		},
		{
			name:     "with class",
			props:    fab.Props{ClassNames: htmx.ClassNames{"bg-red-500": true}},
			expected: `<div class="bg-red-500 fab"></div>`,
		},
		{
			name:     "with multiple classes",
			props:    fab.Props{ClassNames: htmx.ClassNames{"bg-red-500": true, "text-white": true}},
			expected: `<div class="bg-red-500 fab text-white"></div>`,
		},
		{
			name:     "with children",
			props:    fab.Props{},
			children: []htmx.Node{htmx.Text("FAB Content")},
			expected: `<div class="fab">FAB Content</div>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := fab.Fab(
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

func TestFabButton(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    fab.ButtonProps
	}{
		{
			name:     "default",
			props:    fab.ButtonProps{},
			expected: `<div class="btn" role="button" tabindex="0"></div>`,
		},
		{
			name:     "with class",
			props:    fab.ButtonProps{ClassNames: htmx.ClassNames{"bg-red-500": true}},
			expected: `<div class="bg-red-500 btn" role="button" tabindex="0"></div>`,
		},
		{
			name:     "with multiple classes",
			props:    fab.ButtonProps{ClassNames: htmx.ClassNames{"bg-red-500": true, "text-white": true}},
			expected: `<div class="bg-red-500 btn text-white" role="button" tabindex="0"></div>`,
		},
		{
			name:     "with children",
			props:    fab.ButtonProps{},
			children: []htmx.Node{htmx.Text("Button Content")},
			expected: `<div class="btn" role="button" tabindex="0">Button Content</div>`,
		},
		{
			name:     "with tabindex",
			props:    fab.ButtonProps{TabIndex: 1},
			expected: `<div class="btn" role="button" tabindex="1"></div>`,
		},
		{
			name:     "with role",
			props:    fab.ButtonProps{Role: "menu"},
			expected: `<div class="btn" role="menu" tabindex="0"></div>`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := fab.Button(
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
