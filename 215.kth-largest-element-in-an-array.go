/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	l := len(nums) - 1

	// 数组前一半向后做交换
	for i := l >> 1; i >= 0; i-- {
		down(nums, i, l)
	}

	for j := l; j >= 1; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		l--
		down(nums, 0, l)
	}

	return nums[k-1]
}

func down(nums[]int, i, l int) {
	// 最小堆
	min := i

	left, right := i << 1 + 1, i << 1 + 2
	if left <= l && nums[left] < nums[min] {
		min = left
	}

	if right <= l && nums[right] < nums[min] {
		min = right
	}

	if i != min {
		// swap
		nums[i], nums[min] = nums[min], nums[i]
		down(nums, min, l)
	}
}
// @lc code=end

