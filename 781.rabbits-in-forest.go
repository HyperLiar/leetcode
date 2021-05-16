/*
 * @lc app=leetcode id=781 lang=golang
 *
 * [781] Rabbits in Forest
 */

// @lc code=start
func numRabbits(answers []int) int {
    m := make(map[int]int)
	res := 0
	for _, answer := range answers {
		m[answer]++
	}

	for answer, count := range m {
		if answer == 0 {
			res += count
		} else {
			if count % (answer+1) != 0 {
				res += (count / (answer+1) + 1) * (answer+1)
			} else {
				res += count
			}
		}
	}

	return res
}
// @lc code=end

