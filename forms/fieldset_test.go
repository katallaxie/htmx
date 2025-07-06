package forms_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/forms"
	"github.com/stretchr/testify/require"
)

func TestFieldset(t *testing.T) {
	tests := []struct {
		name     string
		props    forms.FieldsetProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    forms.FieldsetProps{},
			expected: "<fieldset class=\"fieldset\"></fieldset>",
			children: nil,
		},
		{
			name:     "with-children",
			props:    forms.FieldsetProps{},
			expected: "<fieldset class=\"fieldset\"><legend>Example Legend</legend><p>Example paragraph.</p></fieldset>",
			children: []htmx.Node{
				htmx.Legend(htmx.Text("Example Legend")),
				htmx.P(htmx.Text("Example paragraph.")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := forms.Fieldset(
				forms.FieldsetProps{},
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}

func TestLegend(t *testing.T) {
	tests := []struct {
		name     string
		props    forms.LegendProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    forms.LegendProps{},
			expected: "<legend class=\"fieldset-legend\"></legend>",
			children: nil,
		},
		{
			name:     "with-children",
			props:    forms.LegendProps{},
			expected: "<legend class=\"fieldset-legend\">Example Legend</legend>",
			children: []htmx.Node{htmx.Text("Example Legend")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := forms.Legend(
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
