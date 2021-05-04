/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start
func nextPermutation(nums []int) {
	i := len(nums) - 2

	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	// i is the min index out of ascending order

	if i >= 0 {
		// find the pos to swap with i
		j := len(nums) - 1

		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	reverse(&nums, i+1, len(nums)-1)
}

func reverse(nums *[]int, i, j int) {
	for i < j {
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
		i++
		j--
	}
}

// @lc code=end

