func subsets(nums []int) [][]int {
    var rs [][]int
    var backtracking func(start int, subset []int)

    backtracking = func(start int, subset []int) {
        // tmp := make([]int, len(subset))
        // copy(tmp, subset)
        // rs = append(rs, tmp)
        rs = append(rs, append([]int{}, subset...))

        for i := start; i < len(nums); i++ {
            subset = append(subset, nums[i])
            backtracking(i+1, subset)
            subset = subset[:len(subset)-1]
        }
    }

    backtracking(0, []int{})
    return rs
}
