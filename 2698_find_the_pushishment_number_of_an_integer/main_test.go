package main

import (
	"strconv"
	"testing"
)

func punishmentNumber(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		square := i * i
		sqStr := strconv.Itoa(square)
		if canPartition(sqStr, i, 0) {
			total += square
		}
	}
	return total
}

func canPartition(numStr string, target int, start int) bool {
	if start == len(numStr) {
		return target == 0
	}
	for end := start; end < len(numStr); end++ {
		part, _ := strconv.Atoi(numStr[start : end+1])
		if part > target {
			break
		}
		if canPartition(numStr, target-part, end+1) {
			return true
		}
	}
	return false
}

func TestPunishmentNumber(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{10, 182},
		{37, 1478},
	}

	for _, tt := range tests {
		result := punishmentNumber(tt.n)
		if result != tt.expected {
			t.Errorf("punishmentNumber(%v ) = %d; want %d", tt.n, result, tt.expected)
		}
	}
}
