package dropdowns_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/dropdowns"
	"github.com/stretchr/testify/require"
)

func TestDropdown(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<details class=\"dropdown\"></details>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<details class=\"custom-class dropdown\"></details>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := dropdowns.Dropdown(
				dropdowns.DropdownProps{
					ClassNames: test.classes,
				},
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}
