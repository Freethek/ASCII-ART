package renderer

import (
	"testing"
)

func TestRender(t *testing.T) {
	bannerMap := map[rune][]string{
		'A': []string{"line1", "line2", "line3", "line4", "line5", "line6", "line7", "line8"},
	}

	//defining the test input into the function and the expected
	test := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"\\n", "\n"},
		{"A", "line1\nline2\nline3\nline4\nline5\nline6\nline7\nline8\n"},
	}

	//looping through and parsing each input to the function
	//and comparin the output wth the expected
	for _, tt := range test {
		//
		result := Render(tt.input, bannerMap)
		if result != tt.expected {
			t.Errorf("input %q: got %q want %q", tt.input, result, tt.expected)
		}
	}

}
