/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	hor, ver := len(grid), len(grid[0])
	sum := 0
	for i := 0; i < hor; i++ {
		for j := 0; j < ver; j++ {
			if grid[i][j] == '1' {
				sum++
				dfs(grid, i, j, hor, ver)
			}
		}
	}

	return sum
}
func dfs(grid [][]byte, i, j, hor, ver int) {
	if i < 0 || j < 0 || i == hor || j == ver {
		return
	}

	if grid[i][j] == '1' {
		grid[i][j] = '0'
	} else {
		return
	}

	dfs(grid, i-1, j, hor, ver)
	dfs(grid, i+1, j, hor, ver)
	dfs(grid, i, j-1, hor, ver)
	dfs(grid, i, j+1, hor, ver)
}

// @lc code=end

