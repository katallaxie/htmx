package tooltips_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/tooltips"
	"github.com/stretchr/testify/require"
)

func TestTooltip(t *testing.T) {
	tests := []struct {
		name     string
		props    tooltips.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    tooltips.Props{},
			expected: "<div class=\"tooltip\" data-tip=\"\"></div>",
			children: nil,
		},
		{
			name:     "open",
			props:    tooltips.Props{Open: true},
			expected: "<div class=\"tooltip tooltip-open\" data-tip=\"\"></div>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := tooltips.Tooltip(
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

func TestPrimary(t *testing.T) {
	tests := []struct {
		name     string
		props    tooltips.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    tooltips.Props{},
			expected: "<div class=\"tooltip tooltip-primary\" data-tip=\"\"></div>",
			children: nil,
		},
		{
			name:     "open",
			props:    tooltips.Props{Open: true},
			expected: "<div class=\"tooltip tooltip-open tooltip-primary\" data-tip=\"\"></div>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := tooltips.Primary(
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

func TestSecondary(t *testing.T) {
	tests := []struct {
		name     string
		props    tooltips.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    tooltips.Props{},
			expected: "<div class=\"tooltip tooltip-secondary\" data-tip=\"\"></div>",
			children: nil,
		},
		{
			name:     "open",
			props:    tooltips.Props{Open: true},
			expected: "<div class=\"tooltip tooltip-open tooltip-secondary\" data-tip=\"\"></div>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := tooltips.Secondary(
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

func TestSuccess(t *testing.T) {
	tests := []struct {
		name     string
		props    tooltips.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    tooltips.Props{},
			expected: "<div class=\"tooltip tooltip-success\" data-tip=\"\"></div>",
			children: nil,
		},
		{
			name:     "open",
			props:    tooltips.Props{Open: true},
			expected: "<div class=\"tooltip tooltip-open tooltip-success\" data-tip=\"\"></div>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := tooltips.Success(
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

func TestWarning(t *testing.T) {
	tests := []struct {
		name     string
		props    tooltips.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    tooltips.Props{},
			expected: "<div class=\"tooltip tooltip-warning\" data-tip=\"\"></div>",
			children: nil,
		},
		{
			name:     "open",
			props:    tooltips.Props{Open: true},
			expected: "<div class=\"tooltip tooltip-open tooltip-warning\" data-tip=\"\"></div>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := tooltips.Warning(
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

func TestError(t *testing.T) {
	tests := []struct {
		name     string
		props    tooltips.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    tooltips.Props{},
			expected: "<div class=\"tooltip tooltip-error\" data-tip=\"\"></div>",
			children: nil,
		},
		{
			name:     "open",
			props:    tooltips.Props{Open: true},
			expected: "<div class=\"tooltip tooltip-error tooltip-open\" data-tip=\"\"></div>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := tooltips.Error(
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
