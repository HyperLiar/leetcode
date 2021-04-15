/*
 * @lc app=leetcode id=507 lang=golang
 *
 * [507] Perfect Number
 */

// @lc code=start
func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	max := num
	sum := 1
    for i := 2; i < max; i++ {
		if num % i == 0 {
			max = num/i
			if max == i {
				sum += i
			} else {
				sum += max + i
			}
		}
	}

	return sum == num
}
// @lc code=end

