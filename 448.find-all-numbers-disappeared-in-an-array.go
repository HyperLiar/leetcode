/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
    for i := 0; i < len(nums); i++ {
		j := abs(nums[i])-1
		if nums[j] > 0 {
			nums[j] *= -1
		}
	}

	res := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}

	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
// @lc code=end

