/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	valueMap := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		valueMap[nums[i]] = i
	}

	for i := 0; i <= len(nums) / 2; i++ {
		bias := target - nums[i]
		if _, exist := valueMap[bias]; exist && valueMap[bias] != i {
			return []int{i, valueMap[bias]}
		}
	}

	return []int{}
}
// @lc code=end

