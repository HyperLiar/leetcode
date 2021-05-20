/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
var dirs = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 || len(board) == 0 {
		return false
	}

	if len(board)*len(board[0]) < len(word) {
		return false
	}

	// search
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {1.
			if dfs(board, word, i, j) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, i, j int) bool {
	if len(word) == 0 {
		return true
	}

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}

	if board[i][j] != word[0] {
		return false
	}

	temp := board[i][j]
	board[i][j] = '#'

	for _, dir := range dirs {
		ii := i + dir[0]
		jj := j + dir[1]

		if dfs(board, word[1:], ii, jj) {
			return true
		}
	}

	board[i][j] = temp
	return false
}

// @lc code=end

