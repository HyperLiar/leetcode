/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		sum := 0

		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}

	return count
}

// @lc code=end

