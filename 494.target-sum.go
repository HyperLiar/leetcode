/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	m := make(map[int]int)
	temp := make(map[int]int)
	m[nums[0]]++
	m[-nums[0]]++
	for i := 1; i < len(nums); i++ {
		for sum, count := range m {
			temp[sum+nums[i]] += count
			temp[sum-nums[i]] += count
		}

		m = map[int]int{}
		for sum, count := range temp {
			m[sum] = count
		}
		temp = map[int]int{}

		fmt.Println(m)
	}

	return m[target]
}

// @lc code=end

