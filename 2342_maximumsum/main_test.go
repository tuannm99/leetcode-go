func maximumSum(nums []int) int {
    // Stores the max number for each digit sum
	digitSumMap := make(map[int]int)
	maxSum := -1

	for _, num := range nums {
		dSum := digitSum(num)

		// If we already have a number with this digit sum, check for max pair sum
		if val, exists := digitSumMap[dSum]; exists {
			if num+val > maxSum {
				maxSum = num + val
			}
			// Store the max number for this digit sum
			if num > val {
				digitSumMap[dSum] = num
			}
		} else {
			// Store the first number for this digit sum
			digitSumMap[dSum] = num
		}
	}

	return maxSum
}


func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}


