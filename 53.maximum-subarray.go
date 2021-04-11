/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
    current, maxValue := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		current = max(nums[i], current + nums[i])
		maxValue = max(maxValue, current)
	}

	return maxValue
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
// @lc code=end

