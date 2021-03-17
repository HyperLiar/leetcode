/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
    bm := make(map[byte]byte)
	bm[')'] = '('
	bm[']'] = '['
	bm['}'] = '{'

	if len(s) & 1 == 1 {
		return false
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if top, ok := bm[s[i]]; ok {
			if len(stack) == 0 || top != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[0:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
// @lc code=end

