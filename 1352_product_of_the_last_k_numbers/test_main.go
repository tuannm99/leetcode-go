package main

import "testing"

type ProductOfNumbers struct {
	prefix []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		prefix: []int{1}, // Prefix product, starting with 1 for convenience
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.prefix = []int{1} // Reset on zero
	} else {
		lastProduct := this.prefix[len(this.prefix)-1]
		this.prefix = append(this.prefix, lastProduct*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.prefix)
	if k >= n {
		return 0 // If k reaches a zero reset point
	}
	return this.prefix[n-1] / this.prefix[n-1-k]

}

func TestProductOfNumbers(t *testing.T) {
	pOfNums := ProductOfNumbers{}

	pOfNums.Add(3)
	pOfNums.Add(0)
	pOfNums.Add(2)
	pOfNums.Add(5)

	pOfNums.Add(4)
	pOfNums.Add(2)
	pOfNums.Add(3)
	pOfNums.Add(4)
	pOfNums.Add(8)
	pOfNums.Add(2)

	tests := []struct {
		k        int
		expected int
	}{
		{1, 2},    // Last number: 2
		{2, 16},   // Last two: 8 * 2
		{3, 64},   // Last three: 4 * 8 * 2
		{4, 768},  // Last four: 3 * 4 * 8 * 2
		{5, 1536}, // Last five: 2 * 3 * 4 * 8 * 2
	}

	for _, test := range tests {
		result := pOfNums.GetProduct(test.k)
		if result != test.expected {
			t.Errorf("GetProduct(%d) = %d; expected %d", test.k, result, test.expected)
		}
	}
}
