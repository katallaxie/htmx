package collapsible_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/collapsible"
	"github.com/stretchr/testify/require"
)

func TestCollapsible(t *testing.T) {
	tests := []struct {
		name     string
		props    collapsible.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    collapsible.Props{},
			expected: "<div class=\"collapse\" tabindex=\"0\"></div>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := collapsible.Collapse(
				collapsible.Props{},
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}
