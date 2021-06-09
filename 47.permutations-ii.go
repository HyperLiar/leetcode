/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	visit := make([]int, len(nums))
	sort.Ints(nums)
	dfs(&res, nums, []int{}, &visit)
	return res
}

func dfs(res *[][]int, nums, cur []int, visit *[]int) {
	if len(cur) == len(nums) {
		*res = append(*res, append([]int{}, cur...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if (*visit)[i] == 1 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && (*visit)[i-1] == 1 {
			continue
		}

		(*visit)[i] = 1
		dfs(res, nums, append(cur, nums[i]), visit)
		(*visit)[i] = 0
	}
}

// @lc code=end

