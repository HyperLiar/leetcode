/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
    m := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		}

		m[nums[i]] = struct{}{}
	}

	return false
}
// @lc code=end

