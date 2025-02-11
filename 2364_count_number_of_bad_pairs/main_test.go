package main

import "testing"

func countBadPairs(nums []int) int64 {
	var goodPairs int64
	count := make(map[int64]int64)

	totalPairs := int64(len(nums)) * (int64(len(nums)) - 1) / 2 // Total pairs

	for i, num := range nums {
		diff := int64(i - num)
		goodPairs += count[diff] // Good pairs
		count[diff]++
	}

	// Bad pairs = Total - Good pairs
	return totalPairs - goodPairs
}

func TestCountBadPairs(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int64
	}{
		{[]int{4, 1, 3, 3}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
	}

	for _, tt := range tests {
		result := countBadPairs(tt.nums)
		if result != tt.expected {
			t.Errorf("countBadPairs(%v) = %d; want %d", tt.nums, result, tt.expected)
		}
	}
}
