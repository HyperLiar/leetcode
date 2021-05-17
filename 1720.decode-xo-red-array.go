/*
 * @lc app=leetcode id=1720 lang=golang
 *
 * [1720] Decode XORed Array
 */

// @lc code=start
func decode(encoded []int, first int) []int {
	temp := first
	for i := 0; i < len(encoded); i++ {
		encoded[i] = first ^ encoded[i]
		first = encoded[i]
	}
	return append([]int{temp}, encoded...)
}

// @lc code=end

