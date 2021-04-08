/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
func findPeakElement(nums []int) int {
    start, end := 0, len(nums)-1

	for start < end-1 {
		mid := (start + end) >> 1

		if mid == 0 || mid == len(nums)-1{
			return mid
		}

		if nums[mid-1] > nums[mid] {
			end = mid
		} else if nums[mid+1] > nums[mid] {
			start = mid
		} else {
			return mid
		}
	}

	if nums[start] > nums[end] {
		return start
	}

	return end
}
// @lc code=end

