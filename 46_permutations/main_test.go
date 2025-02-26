func permute(nums []int) [][]int {
     var rs [][]int
    n := len(nums)

    var backtracking func(permutations []int, used map[int]bool)
    backtracking = func(permutations []int, used map[int]bool) {
        if len(permutations) == n {
            temp := make([]int, len(permutations))
            copy(temp, permutations)
            rs = append(rs, temp)
            return
        }

        for _, num := range nums {
            if used[num] {
                continue
            }
            permutations = append(permutations, num)
            used[num] = true

            backtracking(permutations, used)
            permutations = permutations[:len(permutations)-1]
            delete(used, num)
        }
    }

    backtracking([]int{}, make(map[int]bool))
    return rs
}
