/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
    rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	if cols == 0 {
		return 0
	}

	// build lenMatrix
	lenMatrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		lenMatrix[i] = make([]int, cols)
	}

	for i := rows-1; i >= 0; i-- {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				if i == rows - 1 {
					lenMatrix[i][j] = 1
				} else {
					lenMatrix[i][j] = lenMatrix[i+1][j]+1
				}
			}
		}
	}

	maxRectangle := 0
	// check rectangle
	for i := 0; i < rows; i++ {
		x := largestRectangleArea(lenMatrix[i])
		if x > maxRectangle {
			maxRectangle = x
		}
	}

	return maxRectangle
}

func largestRectangleArea(heights []int) int {
	l := len(heights)

	dpLeft := make([]int, l)
	dpRight := make([]int, l)

	for i := 1; i < l; i++ {
		dpLeft[i] = i
		j := i

		for dpLeft[j]-1 >= 0 && heights[dpLeft[j]-1] >= heights[i] {
			j = dpLeft[j]-1
		}
		dpLeft[i] = dpLeft[j]
	}

	dpRight[l-1] = l-1

	for i := l-2; i >= 0; i-- {
		dpRight[i] = i
		j := i
		for dpRight[j]+1 <= l-1 && heights[dpRight[j]+1] >= heights[i] {
			j = dpRight[j]+1
		}
		dpRight[i] = dpRight[j]
	}

	result := 0
	for i := 0; i < l; i++ {
		x := heights[i] * (dpRight[i] - dpLeft[i] + 1)

		if x > result {
			result = x
		}
	}

	return result
}
// @lc code=end

