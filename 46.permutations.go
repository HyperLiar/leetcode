/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
    result := make([][]int, 0)
	cur := make([]int, 0)

	backTrack(&result, &cur, nums, 0, len(nums))

	return result
}

func backTrack(result *[][]int, cur *[]int, nums []int, now, max int) {
	if len(*cur) == max {
		temp := append([]int{}, *cur...)
		*result = append(*result, temp)
		return
	}

	for i := now; i < max; i++ {
		temp := i
		backTrack()
	}
}
// @lc code=end

