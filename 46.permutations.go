/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
    result := x
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
		nums[now], nums[i] = nums[i], nums[now]
		*cur = append(*cur, nums[now])
		backTrack(result, cur, nums, now+1, max)
		*cur = (*cur)[0:len(*cur)-1]
		nums[now], nums[i] = nums[i], nums[now]
	}
}
// @lc code=end

