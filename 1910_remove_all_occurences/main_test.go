package main

import (
	"testing"
)

func removeOccurrences(s string, part string) string {
	for i := 0; i <= len(s)-len(part); i++ {
		if s[i:i+len(part)] == part {
			s = s[0:i] + s[i+len(part):]
			i = -1
		}
	}
	return s
}

func TestRemoveOccurrences(t *testing.T) {
	tests := []struct {
		s        string
		part     string
		expected string
	}{
		{"daabcbaabcbc", "abc", "dab"},
		{"axxxxyyyyb", "xy", "ab"},
		{"hello", "ll", "heo"},
		{"aaaa", "aa", ""},
		{"abcde", "xyz", "abcde"},
		{"abababa", "aba", "b"},
		{"xyzxyz", "xyz", ""},
	}

	for _, tt := range tests {
		result := removeOccurrences(tt.s, tt.part)
		if result != tt.expected {
			t.Errorf("removeOccurrences(%q, %q) = %q; want %q", tt.s, tt.part, result, tt.expected)
		}
	}
}
