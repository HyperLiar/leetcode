/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	left, cur, right := 0, 0, len(nums) - 1

	for cur <= right {
		if nums[cur] == 0 {
			nums[left], nums[cur] = nums[cur], nums[left]
			left, cur = left+1, cur+1
		} else if nums[cur] == 2 {
			nums[cur], nums[right] = nums[right], nums[cur]
			right--
		} else {
			cur++
		}
	}

	// r, w, b := 0, 0, 0

	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] == 0 {
	// 		r++
	// 	} else if nums[i] == 1 {
	// 		w++
	// 	} else {
	// 		b++
	// 	}
	// }

	// i := 0
	// for i < r {
	// 	nums[i] = 0
	// 	i++
	// }
	// for i < r+w {
	// 	nums[i] = 1
	// 	i++
	// }

	// for i < len(nums) {
	// 	nums[i] = 2
	// 	i++
	// }
}
// @lc code=end

