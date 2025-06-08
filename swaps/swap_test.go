package swaps_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/swaps"
	"github.com/stretchr/testify/require"
)

func TestOn(t *testing.T) {
	tests := []struct {
		name     string
		props    swaps.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    swaps.Props{},
			expected: "<div class=\"swap-on\"></div>",
			children: nil,
		},
		{
			name:     "with class",
			props:    swaps.Props{ClassNames: htmx.ClassNames{"custom-class": true}},
			expected: "<div class=\"custom-class swap-on\"></div>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := swaps.On(
				tt.props,
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}

func TestOff(t *testing.T) {
	tests := []struct {
		name     string
		props    swaps.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    swaps.Props{},
			expected: "<div class=\"swap-off\"></div>",
			children: nil,
		},
		{
			name:     "with class",
			props:    swaps.Props{ClassNames: htmx.ClassNames{"custom-class": true}},
			expected: "<div class=\"custom-class swap-off\"></div>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := swaps.Off(
				tt.props,
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		name     string
		props    swaps.Props
		expected string
		children []htmx.Node
	}{
		{
			name:     "default",
			props:    swaps.Props{},
			expected: "<label class=\"swap\"></label>",
			children: nil,
		},
		{
			name:     "with class",
			props:    swaps.Props{ClassNames: htmx.ClassNames{"custom-class": true}},
			expected: "<label class=\"custom-class swap\"></label>",
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := swaps.Swap(
				tt.props,
				tt.children...,
			)

			var buf bytes.Buffer
			err := component.Render(&buf)
			require.NoError(t, err)

			require.Equal(t, tt.expected, buf.String())
		})
	}
}
