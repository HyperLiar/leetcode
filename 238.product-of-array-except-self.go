/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	productArr := make([]int, len(nums))
	product := 1
	productArr[0] = 1
	
	for i := 1; i < len(nums); i++ {
		productArr[i] = productArr[i-1] * nums[i-1]
	}

	for i := len(nums)-1; i >= 0; i-- {
		productArr[i] *= product
		product *= nums[i]
	}

	return productArr
}

// @lc code=end

