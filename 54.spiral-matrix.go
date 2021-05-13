/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	row, col := len(matrix), len(matrix[0])
	res := make([]int, 0)
	i, j := 0, 0

	for i <= row/2 && j <= col/2 {
		if row != i*2 {
			for k := j; k < col-j; k++ {
				res = append(res, matrix[i][k])
			}
		}
		if col != j*2 {
			for k := i + 1; k < row-i; k++ {
				res = append(res, matrix[k][col-j-1])
			}
		}

		if row-2*i > 1 {
			for k := col - j - 2; k >= j; k-- {
				res = append(res, matrix[row-i-1][k])
			}
		}

		if col-2*j > 1 {
			for k := row - i - 2; k >= i+1; k-- {
				res = append(res, matrix[k][j])
			}
		}
		i, j = i+1, j+1
	}

	return res
}

// @lc code=end

