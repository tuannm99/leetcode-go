package main

import "testing"

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func clearDigits(s string) string {
	var result []rune

	for _, c := range s {
		if isDigit(c) {
			if len(result) > 0 && !isDigit(result[len(result)-1]) {
				result = result[:len(result)-1]
			} else {
				result = append(result, c)
			}

		} else {
			result = append(result, c)
		}
	}

	return string(result)
}

func TestClearDigits(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abc123", ""},
		{"a1b2c3", ""},
		{"123abc456", "123"},
		{"hello42world5", "helworl"},
		{"9abc8", "9ab"},
		{"no123digits", "3digits"},
		{"", ""},
		{"12345", "12345"},
	}

	for _, tt := range tests {
		result := clearDigits(tt.input)
		if result != tt.expected {
			t.Errorf("clearDigits(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
