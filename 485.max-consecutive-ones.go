/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
    maxCount, nowCount := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			nowCount++
		} else {
			nowCount = 0
		}

		if nowCount > maxCount {
			maxCount = nowCount
		}
	}

	return maxCount
}
// @lc code=end

