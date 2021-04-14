/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	result := make([][]int, 1)
	result[0] = []int{}

	for i := 0; i < len(nums); i++ {
		tempResult := make([][]int, 0)
		for j := 0; j < len(result); j++ {
			temp := make([]int, len(result[j]))
			copy(temp, result[j])
			tempResult = append(tempResult, append(temp, nums[i]))
		}
		result = append(result, tempResult...)
	}
	return result
}
// @lc code=end

