/*
 * @lc app=leetcode id=581 lang=golang
 *
 * [581] Shortest Unsorted Continuous Subarray
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
    left, right := 0, len(nums)-1
	for left < len(nums)-1 && nums[left] <= nums[left+1] {
		left++
	}

	for right > left && nums[right-1] <= nums[right] {
		right--
	}

	if left == right {
		return 0
	}

	for i := left+1; i < len(nums); i++ {
		for left >= 0 && nums[i] < nums[left] {
                left--
		}
	}

	for i := right -1; i >=0; i-- {
		for right < len(nums) && nums[i] > nums[right] {
			right++
		}
	}
	return right - left - 1
}
// @lc code=end

