package cards_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/cards"
	"github.com/stretchr/testify/require"
)

func TestCard(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    cards.Props
	}{
		{
			name:     "Card with no children",
			expected: `<div class="card"></div>`,
			children: nil,
			props:    cards.Props{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := cards.Card(
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

func TestCardBorder(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    cards.Props
	}{
		{
			name:     "CardBorder with no children",
			expected: `<div class="card card-border"></div>`,
			children: nil,
			props:    cards.Props{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := cards.CardBorder(
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

func TestBody(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name     string
		expected string
		children []htmx.Node
		props    cards.BodyProps
	}{
		{
			name:     "Body with no children",
			expected: `<div class="card-body"></div>`,
			children: nil,
			props:    cards.BodyProps{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := cards.Body(
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
