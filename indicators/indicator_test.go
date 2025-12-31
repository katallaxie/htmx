package indicators_test

import (
	"bytes"
	"testing"

	"github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/indicators"

	"github.com/stretchr/testify/require"
)

func TestBadge(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    indicators.Props
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<div class=\"indicator\"></div>",
			props:    indicators.Props{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<div class=\"custom-class indicator\"></div>",
			props: indicators.Props{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := indicators.Indicator(
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

func TestItem(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    indicators.ItemProps
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<div class=\"indicator-item\"></div>",
			props:    indicators.ItemProps{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<div class=\"custom-class indicator-item\"></div>",
			props: indicators.ItemProps{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := indicators.Item(
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

func TestItemPrimary(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    indicators.ItemProps
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<div class=\"badge badge-primary indicator-item\"></div>",
			props:    indicators.ItemProps{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<div class=\"badge badge-primary custom-class indicator-item\"></div>",
			props: indicators.ItemProps{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := indicators.BadgePrimary(
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

func TestItemSecondary(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    indicators.ItemProps
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<div class=\"badge badge-secondary indicator-item\"></div>",
			props:    indicators.ItemProps{},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<div class=\"badge badge-secondary custom-class indicator-item\"></div>",
			props: indicators.ItemProps{
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := indicators.BadgeSecondary(
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
