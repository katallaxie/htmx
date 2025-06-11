package typography_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/typography"
	"github.com/stretchr/testify/require"
)

func TestEr(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<span class=\"text-error\"></span>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<span class=\"custom-class text-error\"></span>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<span class=\"text-error\">child</span>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := typography.Error(
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

func TestInfo(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<span class=\"text-info\"></span>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<span class=\"custom-class text-info\"></span>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<span class=\"text-info\">child</span>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := typography.Info(
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

func TestSuccess(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<span class=\"text-success\"></span>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<span class=\"custom-class text-success\"></span>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<span class=\"text-success\">child</span>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := typography.Success(
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

func TestWarning(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<span class=\"text-warning\"></span>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<span class=\"custom-class text-warning\"></span>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<span class=\"text-warning\">child</span>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := typography.Warning(
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

func TestPrimary(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<span class=\"text-primary\"></span>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<span class=\"custom-class text-primary\"></span>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<span class=\"text-primary\">child</span>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := typography.Primary(
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

func TestSecondary(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<span class=\"text-secondary\"></span>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<span class=\"custom-class text-secondary\"></span>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<span class=\"text-secondary\">child</span>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := typography.Secondary(
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
