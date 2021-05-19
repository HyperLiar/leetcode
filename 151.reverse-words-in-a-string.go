/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverse(m []string, i int, j int){
	for {
        if i > j {
			break
		}
		m[i], m[j] = m[j], m[i]
		i++
		j--
	}
}

func reverseWords(s string) string {
	ss := strings.Fields(s)
	reverse(ss,0,len(ss)-1)
    return strings.Join(ss," ")
}
// @lc code=end

