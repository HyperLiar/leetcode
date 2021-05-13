/*
 * @lc app=leetcode id=415 lang=golang
 *
 * [415] Add Strings
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	var sum string
	var remainder, s int
	if len(num2) > len(num1) {
		num1, num2 = num2, num1
	}
	l1, l2 := len(num1), len(num2)

	for i := 0; i < l1; i++ {
		s = int(num1[l1-i-1]-'0') + remainder
		if i < l2 {
			s += int(num2[l2-i-1] - '0')
		}
		if s > 9 {
			s = s - 10
			remainder = 1
		} else {
			remainder = 0
		}
		sum = strconv.Itoa(s) + sum
	}

	if remainder == 1 {
		sum = "1" + sum
	}

	return sum
}

// @lc code=end

