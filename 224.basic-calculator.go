/*
 * @lc app=leetcode id=224 lang=golang
 *
 * [224] Basic Calculator
 */

// @lc code=start
func calculate(s string) int {
	stack := make([]int, 0)
	res, operand, sign := 0, 0, 1

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch-'0' >= 0 && ch-'0' < 10 {
			operand = 10*operand + int(ch-'0')
		} else if ch == '+' {
			res += sign * operand
			operand, sign = 0, 1
		} else if ch == '-' {
			res += sign * operand
			operand, sign = 0, -1
		} else if ch == '(' {
			stack = append(stack, res, sign)
			res, sign = 0, 1
		} else if ch == ')' {
			res += sign * operand
			res *= stack[len(stack)-1]
			res += stack[len(stack)-2]
			stack = stack[0 : len(stack)-2]
			operand = 0
		}
	}

	return res + sign*operand
}

// @lc code=end

