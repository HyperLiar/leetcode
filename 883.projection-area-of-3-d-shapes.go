/*
 * @lc app=leetcode id=883 lang=golang
 *
 * [883] Projection Area of 3D Shapes
 */

// @lc code=start
func projectionArea(grid [][]int) int {
    row, col := len(grid), len(grid[0])
	rowMax := make([]int, row)
	colMax := make([]int, col)
	res := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] != 0 {
				res++
				if grid[i][j] > rowMax[i] {
					rowMax[i] = grid[i][j]
				}

				if grid[i][j] > colMax[j] {
					colMax[j] = grid[i][j]
				}
			}
		}
	}

	for i := 0; i < row; i++ {
		res += rowMax[i]
	}

	for j := 0; j < col; j++ {
		res += colMax[j]
	}
	return res
}
// @lc code=end

