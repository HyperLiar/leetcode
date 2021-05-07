/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */

// @lc code=start
var matrixGlobal [][]int
var targetGlobal int

func searchMatrix(matrix [][]int, target int) bool {
	// col := len(matrix)
	// if col == 0 {
	// 	return false
	// }
	// row := len(matrix[0])
	// if row == 0 {
	// 	return false
	// }

	// matrixGlobal = matrix
	// targetGlobal = target

	// return helper(0, col-1, 0, row-1)

	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, col := 0, len(matrix[0])-1

	for row < len(matrix) && col >= 0 {
		val := matrix[row][col]

		if val > target {
			col--
		} else if val < target {
			row++
		} else {
			return true
		}
	}

	return false
}

func helper(colUp, colDown, rowLeft, rowRight int) bool {
	col, row := len(matrixGlobal)-1, len(matrixGlobal[0])-1
	if colUp > colDown || rowLeft > rowRight {
		return false
	}

	colMid, rowMid := (colUp+colDown)/2, (rowLeft+rowRight)/2
	fmt.Println("col:", colUp, colDown, colMid)
	fmt.Println("row:", rowLeft, rowRight, rowMid)
	//fmt.Println(matrixGlobal[colMid][rowMid], targetGlobal)

	if colMid < 0 || colMid > col || rowMid < 0 || rowMid > row {
		return false
	}

	if matrixGlobal[colMid][rowMid] > targetGlobal {
		r1 := helper(colUp, colMid-1, rowLeft, rowRight)
		r2 := helper(colMid, colDown, rowLeft, rowMid-1)

		return r1 || r2
	} else if matrixGlobal[colMid][rowMid] < targetGlobal {
		r1 := helper(colUp, colMid, rowMid+1, rowRight)
		r2 := helper(colMid+1, colDown, rowLeft, rowRight)

		return r1 || r2
	} else {
		return true
	}
}

// @lc code=end

