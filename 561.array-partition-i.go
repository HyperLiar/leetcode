/*
 * @lc app=leetcode id=561 lang=golang
 *
 * [561] Array Partition I
 */

// @lc code=start
func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	res := 0
	for i := 0; i <= len(nums)/2-1; i++ {
		res += nums[i*2]
	}

	return res
}

// @lc code=end

