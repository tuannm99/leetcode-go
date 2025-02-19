package main

import (
	"container/heap"
	"testing"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Min Heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// minOperations calculates the minimum operations needed
func minOperations(nums []int, k int) int {
	h := MinHeap(nums)
	heap.Init(&h)

	operations := 0

	// Continue until all elements are >= k
	for h.Len() > 1 && h[0] < k {
		x := heap.Pop(&h).(int) // Smallest
		y := heap.Pop(&h).(int) // Second smallest

		// Create new value
		newValue := min(x, y)*2 + max(x, y)
		heap.Push(&h, newValue)
		operations++
	}

	// If the smallest element is still < k, it's impossible
	if h[0] < k {
		return -1
	}

	return operations
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestMinOperations(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 6, 3},
		{[]int{5, 7, 8, 10}, 12, 2},
		{[]int{1, 1, 1, 1}, 10, -1},
		{[]int{1, 2, 3, 4, 5}, 10, 3},
		{[]int{1, 1, 1}, 3, 2},
	}

	for _, tt := range tests {
		result := minOperations(tt.nums, tt.k)
		if result != tt.expected {
			t.Errorf("minOperations(%v, %d) = %d; want %d", tt.nums, tt.k, result, tt.expected)
		}
	}
}

