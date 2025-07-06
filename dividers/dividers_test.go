package dividers_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/dividers"
	"github.com/stretchr/testify/require"
)

func TestDivider(t *testing.T) {
	tests := []struct {
		name     string
		props    dividers.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    dividers.Props{},
			expected: "<div class=\"divider\"></div>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := dividers.Divider(
				dividers.Props{},
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}

func TestDividerNeutral(t *testing.T) {
	tests := []struct {
		name     string
		props    dividers.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    dividers.Props{},
			expected: "<div class=\"divider divider-neutral\"></div>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := dividers.Neutral(
				dividers.Props{},
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}

func TestDividerPrimary(t *testing.T) {
	tests := []struct {
		name     string
		props    dividers.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    dividers.Props{},
			expected: "<div class=\"divider divider-primary\"></div>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := dividers.Primary(
				dividers.Props{},
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}

func TestDividerSecondary(t *testing.T) {
	tests := []struct {
		name     string
		props    dividers.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    dividers.Props{},
			expected: "<div class=\"divider divider-secondary\"></div>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := dividers.Secondary(
				dividers.Props{},
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}

func TestDividerSuccess(t *testing.T) {
	tests := []struct {
		name     string
		props    dividers.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    dividers.Props{},
			expected: "<div class=\"divider divider-success\"></div>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := dividers.Success(
				dividers.Props{},
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}
