/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
func findMin(nums []int) int {
    left, right := 0, len(nums)-1

    for left < right {
        mid := (left+right) / 2
        if nums[mid] < nums[right] {
            right = mid
        } else if nums[mid] > nums[right] {
            left, right = mid+1, right
        }
    }

    return nums[left]
}
// @lc code=end

