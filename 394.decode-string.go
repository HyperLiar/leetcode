/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func decodeString(s string) string {
	res := ""
	left, right, count := 0, len(s)-1, -1
	
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])

		if s[i] == '[' {
			left++
		} else if s[i] == ']' {
			temp := ""
			for j := len(stack)-1; j >= 0; j-- {
				if stack[j] == '[' {
					res = 
				}
			}
		}
	}
}

func isNum(a byte) int {
	return int(a)
}

// @lc code=end

