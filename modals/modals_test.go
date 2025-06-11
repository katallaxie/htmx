package modals_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/modals"
	"github.com/stretchr/testify/require"
)

func TestModal(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<dialog class=\"modal\" id=\"\"></dialog>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<dialog class=\"custom-class modal\" id=\"\"></dialog>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<dialog class=\"modal\" id=\"\">child</dialog>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := modals.Modal(
				modals.Props{
					ClassNames: test.classes,
				},
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}
