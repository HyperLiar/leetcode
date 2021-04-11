/*
 * @lc app=leetcode id=443 lang=golang
 *
 * [443] String Compression
 */

// @lc code=start
func compress(chars []byte) int {
	charIdx, write := 0, 0

	for read := 0; read < len(chars); read++ {
		if read+1 == len(chars) || chars[read+1] != chars[read] {
			chars[write] = chars[read]
			write++

			// count of char at charIdx is not only 1
			if read > charIdx {
				for _,c := range []byte(strconv.Itoa(read - charIdx + 1)) {
					chars[write] = c
					write++
				}
			}

			charIdx = read+1
		}
	}

	return write
}
// @lc code=end

