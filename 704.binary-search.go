/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func search(nums []int, target int) int {
    start, end := 0, len(nums)-1

	for start <= end {
		pivot := (end+start) / 2

		if nums[pivot] < target {
			start = pivot+1
		} else if nums[pivot] > target {
			end = pivot-1
		} else {
			return pivot
		}
	}

	return  -1
}
// @lc code=end

