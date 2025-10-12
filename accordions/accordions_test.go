package accordions_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/accordions"
	"github.com/stretchr/testify/require"
)

func TestAccordion(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		expected string
		checked  bool
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"collapse\"><input type=\"radio\" name=\"\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"collapse custom-class\"><input type=\"radio\" name=\"\"></div>",
		},
		{
			name:     "checked",
			classes:  nil,
			expected: "<div class=\"collapse\"><input type=\"radio\" name=\"\" checked=\"checked\"></div>",
			checked:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := accordions.Accordion(
				accordions.Props{
					ClassNames: test.classes,
					Checked:    test.checked,
				},
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}

func TestTitle(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"collapse-title\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"collapse-title custom-class\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			expected: "<div class=\"collapse-title\">child</div>",
			children: []htmx.Node{htmx.Text("child")},
		},
		{
			name:     "with classes and children",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"collapse-title custom-class\">child</div>",
			children: []htmx.Node{htmx.Text("child")},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := accordions.Title(
				accordions.TitleProps{
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

func TestContent(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"collapse-content\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"collapse-content custom-class\"></div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := accordions.Content(
				accordions.ContentProps{
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
