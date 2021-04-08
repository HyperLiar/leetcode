/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	rowMax := len(matrix[0]) - 1
    rowStart, rowEnd := 0, rowMax
	colStart, colEnd := 0, len(matrix)-1

	for colStart <= colEnd {
		colMid := (colStart + colEnd) >> 1

		if matrix[colMid][0] > target {
			colEnd = colMid-1
		} else if matrix[colMid][rowMax] < target {
			colStart = colMid+1
		} else {
			colStart = colMid
			break
		}
	}
	if colStart > colEnd {
		return false
	}

	for rowStart <= rowEnd {
		rowMid := (rowStart + rowEnd) >> 1
	 
		if target > matrix[colStart][rowMid] {
			rowStart = rowMid+1
		} else if target < matrix[colStart][rowMid] {
			rowEnd = rowMid-1
		} else {
			return true
		}
	}

	return false
}
// @lc code=end

