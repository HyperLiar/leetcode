/*
 * @lc app=leetcode id=1217 lang=golang
 *
 * [1217] Minimum Cost to Move Chips to The Same Position
 */

// @lc code=start
func minCostToMoveChips(position []int) int {
    odd, even := 0, 0

	for _, val := range position {
		if val & 1 == 1 {
			odd++
		} else {
			even++
		}
	}

	if odd > even {
		return even
	}

	return odd
}
// @lc code=end

