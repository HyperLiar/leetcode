/*
 * @lc app=leetcode id=399 lang=golang
 *
 * [399] Evaluate Division
 */

// @lc code=start
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    m := make(map[string]string)
	for i := 0; i < len(equations); i++ {
		equation := equations[i]
		m[equation[0]] = equation[1]
	}

	for i := 0; i < len(queries); i++ {
		mul := -1.0000

		if queries[0] == queries[1] {
			mul = 1.0000
		} else {
			for 
		}
	}
}
// @lc code=end

