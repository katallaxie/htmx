package badges_test

import (
	"bytes"
	"testing"

	"github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/badges"

	"github.com/stretchr/testify/require"
)

func TestBadge(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    badges.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<span class=\"badge\"></span>",
			props:    badges.Props{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<span class=\"badge custom-class\"></span>",
			props: badges.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := badges.Badge(
				test.props,
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}

func TestNeutral(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    badges.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<span class=\"badge badge-neutral\"></span>",
			props:    badges.Props{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<span class=\"badge badge-neutral custom-class\"></span>",
			props: badges.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := badges.Neutral(
				test.props,
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
		expected string
		props    badges.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<span class=\"badge badge-primary\"></span>",
			props:    badges.Props{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<span class=\"badge badge-primary custom-class\"></span>",
			props: badges.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := badges.Primary(
				test.props,
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
		expected string
		props    badges.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<span class=\"badge badge-secondary\"></span>",
			props:    badges.Props{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<span class=\"badge badge-secondary custom-class\"></span>",
			props: badges.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := badges.Secondary(
				test.props,
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}

func TestAccent(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    badges.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<span class=\"badge badge-accent\"></span>",
			props:    badges.Props{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<span class=\"badge badge-accent custom-class\"></span>",
			props: badges.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := badges.Accent(
				test.props,
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}

func TestGhost(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    badges.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<span class=\"badge badge-ghost\"></span>",
			props:    badges.Props{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<span class=\"badge badge-ghost custom-class\"></span>",
			props: badges.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := badges.Ghost(
				test.props,
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
		expected string
		props    badges.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<span class=\"badge badge-info\"></span>",
			props:    badges.Props{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<span class=\"badge badge-info custom-class\"></span>",
			props: badges.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := badges.Info(
				test.props,
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
		expected string
		props    badges.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<span class=\"badge badge-warning\"></span>",
			props:    badges.Props{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<span class=\"badge badge-warning custom-class\"></span>",
			props: badges.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := badges.Warning(
				test.props,
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
		expected string
		props    badges.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<span class=\"badge badge-error\"></span>",
			props:    badges.Props{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<span class=\"badge badge-error custom-class\"></span>",
			props: badges.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := badges.Error(
				test.props,
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
		expected string
		props    badges.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<span class=\"badge badge-success\"></span>",
			props:    badges.Props{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<span class=\"badge badge-success custom-class\"></span>",
			props: badges.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := badges.Success(
				test.props,
				test.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, test.expected, buf.String())
		})
	}
}
