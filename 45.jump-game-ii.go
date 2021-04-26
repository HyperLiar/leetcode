/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j >= len(nums) {
				break
			}
			dp[i+j] = min(dp[i+j], dp[i]+1)
		}
	}

	//fmt.Println(dp)
	return dp[len(nums)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

// @lc code=end

