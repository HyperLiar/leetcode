/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
    // reach := 0
	// for i := 0; i < len(nums); i++ {
	// 	if i > reach || reach >= len(nums)-1 {
	// 		break
	// 	}
	// 	reach = max(reach, i+nums[i])
	// }

	// return reach >= len(nums)-1

	last := len(nums)-1
	for i := last; i >= 0; i-- {
		if i + nums[i] >= last {
			last = i
		}
	}

	return last == 0
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
// @lc code=end

