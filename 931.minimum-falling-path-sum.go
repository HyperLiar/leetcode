/*
 * @lc app=leetcode id=931 lang=golang
 *
 * [931] Minimum Falling Path Sum
 */

// @lc code=start
func minFallingPathSum(matrix [][]int) int {
	// for i := 1; i < len(matrix); i++ {
	// 	for j := 0; j < len(matrix); j++ {
	// 		min := matrix[i-1][j]

	// 		for offset := -1; offset <= 1; offset += 2 {
	// 			if j+offset >= 0 && j + offset < len(matrix) {
	// 				min = minInt(min, matrix[i-1][j+offset])
	// 			}
	// 		}

	// 		matrix[i][j] += min
	// 	}
	// }

	// return minInt(matrix[len(matrix)-1]...)

	l := len(matrix)
	if l == 1 {
		return matrix[0][0]
	}

	for i := l - 2; i >= 0; i-- {
		matrix[i][0] = matrix[i][0] + minInt(matrix[i+1][0], matrix[i+1][1])
		matrix[i][l-1] = matrix[i][l-1] + minInt(matrix[i+1][l-1], matrix[i+1][l-2])

		for j := 1; j < l-1; j++ {
			dp[i][j] = matrix[i][j] + minInt(matrix[i+1][j-1], matrix[i+1][j], dp[i+1][j+1])
		}
	}

	return minInt(matrix[0]...)
}

func minInt(nums ...int) int {
	min := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	return min
}

// @lc code=end

