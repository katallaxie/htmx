package themes_test

import (
	"bytes"
	"testing"

	"github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/themes"
	"github.com/stretchr/testify/require"
)

func TestTheme(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		theme    string
	}{
		{
			name:     "default theme",
			expected: `<html data-theme="light"></html>`,
			theme:    "light",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := htmx.HTML(
				themes.Theme(test.theme),
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}
