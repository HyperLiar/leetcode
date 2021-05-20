/*
 * @lc app=leetcode id=212 lang=golang
 *
 * [212] Word Search II
 */

// @lc code=start
func findWords(board [][]byte, words []string) []string {
	row, col := len(board), len(board[0])
	trie := build(words)
	res := make([]string, 0)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			dfs(board, i, j, row, col, trie, &res)
		}
	}

	return res
}

func dfs(board [][]byte, i, j, row, col int, trie *Trie, res *[]string) {
	if trie == nil || i < 0 || i >= row || j < 0 || j >= col {
		return
	}

	c := board[i][j]
	if c == '#' || trie.next[c-'a'] == nil {
		return
	}
	trie = trie.next[c-'a']
	if len(trie.word) > 0 {
		*res = append(*res, trie.word)
		trie.word = ""
	}
	board[i][j] = '#'
	dfs(board, i+1, j, row, col, trie, res)
	dfs(board, i-1, j, row, col, trie, res)
	dfs(board, i, j+1, row, col, trie, res)
	dfs(board, i, j-1, row, col, trie, res)
	board[i][j] = c
}

type Trie struct {
	next [26]*Trie
	word string
}

func build(words []string) *Trie {
	root := &Trie{}

	for _, word := range words {
		temp := root
		for i := 0; i < len(word); i++ {
			idx := int(word[i] - 'a')
			if temp.next[idx] == nil {
				temp.next[idx] = &Trie{}
			}
			temp = temp.next[idx]
		}
		temp.word = word
	}

	return root
}

// @lc code=end

