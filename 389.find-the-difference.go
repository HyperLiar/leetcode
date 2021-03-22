/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
    m := make(map[byte]int)

	for _, b := range []byte(s) {
		m[b]++
	}

	for _, b := range []byte(t) {
		m[b]--

		if m[b] < 0 {
			return b
		}
	}

	return byte(0)
}
// @lc code=end

