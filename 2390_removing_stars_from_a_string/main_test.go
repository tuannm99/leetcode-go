package main

import "testing"

func removeStars(s string) string {
	var stack []byte

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}

func TestRemoveStarts(t *testing.T) {

	tests := []struct {
		s        string
		expected string
	}{
		{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
	}

	for _, tt := range tests {
		result := removeStars(tt.s)
		if result != tt.expected {
			t.Errorf("removeStars(%s ) = %s; want %s", tt.s, result, tt.expected)
		}
	}
}
