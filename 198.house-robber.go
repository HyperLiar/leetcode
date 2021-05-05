/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
    dpRobLast := make([]int, len(nums))
	dpNotRobLast := make([]int, len(nums))

	dpRobLast[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dpRobLast[i] = dpNotRobLast[i-1] + nums[i]
		dpNotRobLast[i] = max(dpRobLast[i-1], dpNotRobLast[i-1])
	}

	return max(dpRobLast[len(nums)-1], dpNotRobLast[len(nums)-1])
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
// @lc code=end

