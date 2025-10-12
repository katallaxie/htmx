package lists_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/lists"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	tests := []struct {
		name     string
		props    lists.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    lists.Props{},
			expected: "<ul class=\"list\"></ul>",
		},
		{
			name:     "with class",
			props:    lists.Props{ClassNames: htmx.ClassNames{"custom-class": true}},
			expected: "<ul class=\"custom-class list\"></ul>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := lists.List(
				tt.props,
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}

func TestListRow(t *testing.T) {
	tests := []struct {
		name     string
		props    lists.RowProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    lists.RowProps{},
			expected: "<li class=\"list-row\"></li>",
		},
		{
			name:     "with class",
			props:    lists.RowProps{ClassNames: htmx.ClassNames{"custom-class": true}},
			expected: "<li class=\"custom-class list-row\"></li>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := lists.Row(
				tt.props,
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}

func TestListTitle(t *testing.T) {
	tests := []struct {
		name     string
		props    lists.TitleProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    lists.TitleProps{},
			expected: "<li class=\"opacity-60 p-4 pb-2 text-xs tracking-wide\"></li>",
		},
		{
			name:     "with class",
			props:    lists.TitleProps{ClassNames: htmx.ClassNames{"custom-class": true}},
			expected: "<li class=\"custom-class opacity-60 p-4 pb-2 text-xs tracking-wide\"></li>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := lists.Title(
				tt.props,
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}
