/*
 * @lc app=leetcode id=1190 lang=golang
 *
 * [1190] Reverse Substrings Between Each Pair of Parentheses
 */

// @lc code=start
func reverseParentheses(s string) string {
	if len(s) == 0 {
		return s
	}
	stack := make([]string, 1)
	stack[0] = ""

	for _, b := range s {
		if b == '(' {
			stack = append(stack, "")
		} else if b == ')' {
			r := reverse(stack[len(stack)-1])
			if len(stack) <= 1 {
				stack[0] = r
			} else {
				stack[len(stack)-2] += r
				stack = stack[0 : len(stack)-1]
			}
		} else {
			stack[len(stack)-1] += string(b)
		}
	}

	res := ""
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}

	return res
}

func reverse(s string) string {
	r := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		r[i], r[len(s)-i-1] = r[len(s)-i-1], r[i]
	}

	return string(r)
}

// @lc code=end

