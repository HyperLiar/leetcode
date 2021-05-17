/*
 * @lc app=leetcode id=554 lang=golang
 *
 * [554] Brick Wall
 */

// @lc code=start
func leastBricks(wall [][]int) int {
	var width int
	l := len(wall)
	for i := 0; i < len(wall[0]); i++ {
		width += wall[0][i]
	}
	if width == 1 {
		return l
	}
	brick := map[int]int{}
	for i := 0; i < len(wall); i++ {
		brick[wall[i][0]]++
		for j := 1; j < len(wall[i]); j++ {
			wall[i][j] += wall[i][j-1]
			brick[wall[i][j]]++
		}
	}

	max := 0
	for i, c := range brick {
		if c > max && i != width {
			max = c
		}
	}
	
	return l - max
}

// @lc code=end

