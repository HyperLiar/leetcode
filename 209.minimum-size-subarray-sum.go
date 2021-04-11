/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
    left, minLen := 0, len(nums) + 1

	for i := 0; i < len(nums); i++ {
		target -= nums[i]

		for target <= 0 {
			minLen = min(minLen, i - left + 1)
			target += nums[left]
			left++
		}
	}

	if minLen == len(nums) + 1 {
		return 0
	}

	return minLen
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
// @lc code=end

