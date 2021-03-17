/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
    left, right, l := 0, 0, len(nums)

	for right < l {
		if nums[right] != val {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}

		right++
	}

	return left
}
// @lc code=end

