package alerts_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/alerts"
	"github.com/stretchr/testify/require"
)

func TestAlert(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"alert\" role=\"alert\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"alert custom-class\" role=\"alert\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"alert\" role=\"alert\">child</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := alerts.Alert(
				alerts.Props{
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
			expected: "<div class=\"alert alert-info\" role=\"alert\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"alert alert-info custom-class\" role=\"alert\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"alert alert-info\" role=\"alert\">child</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := alerts.Info(
				alerts.Props{
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
			expected: "<div class=\"alert alert-success\" role=\"alert\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"alert alert-success custom-class\" role=\"alert\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"alert alert-success\" role=\"alert\">child</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := alerts.Success(
				alerts.Props{
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
			expected: "<div class=\"alert alert-warning\" role=\"alert\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"alert alert-warning custom-class\" role=\"alert\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"alert alert-warning\" role=\"alert\">child</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := alerts.Warning(
				alerts.Props{
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

func TestError(t *testing.T) {
	tests := []struct {
		name     string
		classes  htmx.ClassNames
		children []htmx.Node
		expected string
	}{
		{
			name:     "default",
			classes:  nil,
			expected: "<div class=\"alert alert-error\" role=\"alert\"></div>",
		},
		{
			name:     "with classes",
			classes:  htmx.ClassNames{"custom-class": true},
			expected: "<div class=\"alert alert-error custom-class\" role=\"alert\"></div>",
		},
		{
			name:     "with children",
			classes:  nil,
			children: []htmx.Node{htmx.Text("child")},
			expected: "<div class=\"alert alert-error\" role=\"alert\">child</div>",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := alerts.Error(
				alerts.Props{
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
