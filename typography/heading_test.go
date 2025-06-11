package typography_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/typography"
	"github.com/stretchr/testify/require"
)

func TestH1(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<h1 class=\"font-bold text-4xl\"></h1>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<h1 class=\"custom-class font-bold text-4xl\"></h1>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := typography.H1(
				typography.Props{
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

func TestH2(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<h2 class=\"font-bold text-3xl\"></h2>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<h2 class=\"custom-class font-bold text-3xl\"></h2>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := typography.H2(
				typography.Props{
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

func TestH3(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<h3 class=\"font-bold text-2xl\"></h3>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<h3 class=\"custom-class font-bold text-2xl\"></h3>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := typography.H3(
				typography.Props{
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

func TestH4(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<h4 class=\"font-bold text-xl\"></h4>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<h4 class=\"custom-class font-bold text-xl\"></h4>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := typography.H4(
				typography.Props{
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

func TestH5(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<h5 class=\"font-bold text-lg\"></h5>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<h5 class=\"custom-class font-bold text-lg\"></h5>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := typography.H5(
				typography.Props{
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

func TestH6(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<h6 class=\"font-bold text-base\"></h6>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<h6 class=\"custom-class font-bold text-base\"></h6>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := typography.H6(
				typography.Props{
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
