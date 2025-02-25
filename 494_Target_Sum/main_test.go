package main

func findTargetSumWays(nums []int, target int) int {
    count := 0
    backtracking(nums, 0, 0, &count, target)
    return count
}


func backtracking(nums []int, n, sum int, count *int, target int) {
    if len(nums) == n {
        if sum == target {
            *count += 1
        }
        return
    }

    for _, c := range []rune{'-', '+'} {
        if c == '-' {
            sum -= nums[n]
            backtracking(nums, n+1, sum, count, target)
            sum += nums[n]
        } else if c == '+' {
            sum += nums[n]
            backtracking(nums, n+1, sum, count, target)
            sum -= nums[n]
        }
    }
}
