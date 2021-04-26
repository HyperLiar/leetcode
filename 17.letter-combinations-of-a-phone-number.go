/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := make(map[byte]string)
	m['2'] = "abc"
	m['3'] = "def"
	m['4'] = "ghi"
	m['5'] = "jkl"
	m['6'] = "mno"
	m['7'] = "pqrs"
	m['8'] = "tuv"
	m['9'] = "wxyz"

	res := make([]string, 0)
	cur := ""
	backTrack(&res, m, &cur, digits, 0)
	return res

	// res := make([]string, 1)
	// res[0] = ""

	// for i := 0; i < len(digits); i++ {
	// 	t := make([]string, 0)
	// 	idx := digits[i]
	// 	//fmt.Println(digits[i], idx, m[idx])

	// 	for j := 0; j < len(m[idx]); j++ {
	// 		for _, s := range res {
	// 			t = append(t, s+string(m[idx][j]))
	// 		}
	// 	}

	// 	res = t
	// }

	// return res
}

func backTrack(res *[]string, m map[byte]string, cur *string, digits string, idx int) {
	if len(*cur) == len(digits) {
		temp := *cur
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(m[digits[idx]]); i++ {
		*cur = *cur + string(m[digits[idx]][i])
		backTrack(res, m, cur, digits, idx+1)
		*cur = (*cur)[0 : len(*cur)-1]
	}
}

// @lc code=end

