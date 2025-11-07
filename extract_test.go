package main

import "testing"

func TestGetH1FromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputHTML string
		expected  string
	}{
		{
			name:      "valid H1",
			inputHTML: "<h1>valid<h1>",
			expected:  "valid",
		},
		{
			name:      "h3",
			inputHTML: "<h3>this is an h3<h3>",
			expected:  "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1FromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %s; actual: %s", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}

func TestGetPFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputHTML string
		expected  string
	}{
		{
			name:      "valid p",
			inputHTML: "<p>valid<p>",
			expected:  "valid",
		},
		{
			name:      "output",
			inputHTML: "<output>this is an h3<output>",
			expected:  "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected: %s; actual: %s", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}
