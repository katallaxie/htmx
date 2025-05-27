package forms_test

import (
	"bytes"
	"testing"

	htmx "github.com/katallaxie/htmx"
	"github.com/katallaxie/htmx/forms"

	"github.com/stretchr/testify/require"
)

func TestCheckbox(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		props    forms.CheckboxProps
		children []htmx.Node
	}{
		{
			name:     "default",
			expected: "<input class=\"checkbox\" type=\"checkbox\" name=\"test\" value=\"value\">",
			props: forms.CheckboxProps{
				Name:  "test",
				Value: "value",
			},
			children: nil,
		},
		{
			name:     "checked",
			expected: "<input class=\"checkbox\" type=\"checkbox\" name=\"test\" value=\"value\" checked=\"checked\">",
			props: forms.CheckboxProps{
				Name:    "test",
				Value:   "value",
				Checked: true,
			},
			children: nil,
		},
		{
			name:     "disabled",
			expected: "<input class=\"checkbox\" type=\"checkbox\" name=\"test\" value=\"value\" disabled=\"disabled\">",
			props: forms.CheckboxProps{
				Name:     "test",
				Value:    "value",
				Disabled: true,
			},
			children: nil,
		},
		{
			name:     "with classes",
			expected: "<input class=\"checkbox custom-class\" type=\"checkbox\" name=\"test\" value=\"value\">",
			props: forms.CheckboxProps{
				Name:       "test",
				Value:      "value",
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: nil,
		},
		{
			name:     "with children",
			expected: "<input class=\"checkbox custom-class\" type=\"checkbox\" name=\"test\" value=\"value\">",
			props: forms.CheckboxProps{
				Name:       "test",
				Value:      "value",
				ClassNames: htmx.ClassNames{"custom-class": true},
			},
			children: []htmx.Node{htmx.Text("child")},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			component := forms.Checkbox(
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
