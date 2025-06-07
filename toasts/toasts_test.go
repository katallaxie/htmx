package toasts_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/toasts"
	"github.com/stretchr/testify/require"
)

func TestToast(t *testing.T) {
	tests := []struct {
		name     string
		props    toasts.ToastProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default toast",
			props:    toasts.ToastProps{},
			expected: `<div class="toast"></div>`,
			children: nil,
		},
		{
			name: "toast with class",
			props: toasts.ToastProps{
				ClassNames: htmx.ClassNames{"custom-toast": true},
			},
			expected: `<div class="custom-toast toast"></div>`,
			children: nil,
		},
		{
			name:     "toast with children",
			props:    toasts.ToastProps{},
			expected: `<div class="toast">Item 1</div>`,
			children: []htmx.Node{
				htmx.Text("Item 1"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := toasts.Toat(
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

func TestToastTop(t *testing.T) {
	tests := []struct {
		name     string
		props    toasts.ToastProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default toast top",
			props:    toasts.ToastProps{},
			expected: `<div class="toast toast-top"></div>`,
			children: nil,
		},
		{
			name: "toast top with class",
			props: toasts.ToastProps{
				ClassNames: htmx.ClassNames{"custom-toast-top": true},
			},
			expected: `<div class="custom-toast-top toast toast-top"></div>`,
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := toasts.ToastTop(
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

func TestToastBottom(t *testing.T) {
	tests := []struct {
		name     string
		props    toasts.ToastProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default toast bottom",
			props:    toasts.ToastProps{},
			expected: `<div class="toast toast-bottom"></div>`,
			children: nil,
		},
		{
			name: "toast bottom with class",
			props: toasts.ToastProps{
				ClassNames: htmx.ClassNames{"custom-toast-bottom": true},
			},
			expected: `<div class="custom-toast-bottom toast toast-bottom"></div>`,
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := toasts.ToastBottom(
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

func TestToastStart(t *testing.T) {
	tests := []struct {
		name     string
		props    toasts.ToastProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default toast start",
			props:    toasts.ToastProps{},
			expected: `<div class="toast toast-start"></div>`,
			children: nil,
		},
		{
			name: "toast start with class",
			props: toasts.ToastProps{
				ClassNames: htmx.ClassNames{"custom-toast-start": true},
			},
			expected: `<div class="custom-toast-start toast toast-start"></div>`,
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := toasts.ToastStart(
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

func TestToastEnd(t *testing.T) {
	tests := []struct {
		name     string
		props    toasts.ToastProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default toast end",
			props:    toasts.ToastProps{},
			expected: `<div class="toast toast-end"></div>`,
			children: nil,
		},
		{
			name: "toast end with class",
			props: toasts.ToastProps{
				ClassNames: htmx.ClassNames{"custom-toast-end": true},
			},
			expected: `<div class="custom-toast-end toast toast-end"></div>`,
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := toasts.ToastEnd(
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

func TestToastMiddle(t *testing.T) {
	tests := []struct {
		name     string
		props    toasts.ToastProps
		expected string
		children []htmx.Node
	}{
		{
			name:     "default toast middle",
			props:    toasts.ToastProps{},
			expected: `<div class="toast toast-middle"></div>`,
			children: nil,
		},
		{
			name: "toast middle with class",
			props: toasts.ToastProps{
				ClassNames: htmx.ClassNames{"custom-toast-middle": true},
			},
			expected: `<div class="custom-toast-middle toast toast-middle"></div>`,
			children: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := toasts.ToastMiddle(
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
