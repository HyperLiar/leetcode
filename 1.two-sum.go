/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	valueMap := map[int]int{}
	res := []int{}

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if idx, ok := valueMap[diff]; ok {
			res = append(res, i, idx)
			break
		}
		valueMap[nums[i]] = i
	}

	return res
}

// @lc code=end

