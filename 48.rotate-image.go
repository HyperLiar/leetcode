/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l/2; i++ {
		for j := i; j <  l - i - 1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[l-1-j][i]
			matrix[l-1-j][i] = matrix[l-i-1][l-1-j]
			matrix[l-i-1][l-1-j] = matrix[j][l-i-1]
			matrix[j][l-i-1] = temp
		}
	}
}

// @lc code=end

