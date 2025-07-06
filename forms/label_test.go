package forms_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/forms"
	"github.com/stretchr/testify/require"
)

func TestLabel(t *testing.T) {
	tests := []struct {
		name     string
		props    forms.LabelProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    forms.LabelProps{},
			expected: "<label class=\"label\"></label>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := forms.Label(
				forms.LabelProps{},
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}

func TestFloatingLabel(t *testing.T) {
	tests := []struct {
		name     string
		props    forms.LabelProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    forms.LabelProps{},
			expected: "<label class=\"floating-label\"></label>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := forms.FloatingLabel(
				forms.LabelProps{},
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}
