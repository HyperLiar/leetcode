/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l := 1
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		pos := bS(dp, 0, l, nums[i])
		dp[pos] = nums[i]
		if pos == l {
			l++
		}
	}

	return l
}

func bS(nums []int, start, end, num int) int {
	for start < end {
		mid := (start + end) >> 1
		if nums[mid] < num {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return start
}

// @lc code=end

