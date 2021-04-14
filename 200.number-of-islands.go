/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
    hor, ver := len(grid), len(grid[0])

	sum := 0
	for x := 0; x < ver; x++ {
		for y := 0; y < hor; y++ {
			if grid[y][x] == '1' {
				sum++
				dfs(grid, x, y, ver, hor)
			}
		}
	}

	return sum
}

func dfs(grid [][]byte, x, y, ver, hor int) {	
	if x < 0 || y < 0 || x > ver-1 || y > hor-1 {
		return
	}

	if grid[y][x] == '1' {
		grid[y][x] = '0'
	} else {
		return
	}

	dfs(grid, x+1, y, ver, hor)
	dfs(grid, x, y+1, ver, hor)
	dfs(grid, x-1, y, ver, hor)
	dfs(grid, x, y-1, ver, hor)
}
// @lc code=end

