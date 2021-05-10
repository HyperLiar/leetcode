/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	need, have := [26]int{}, [26]int{}
	res := []int{}

	for i := 0; i < len(p); i++ {
		need[p[i]-'a']++
		have[s[i]-'a']++
	}

	var tail int
	for i := len(p); i < len(s); i++ {
		if have == need {
			res = append(res, tail)
		}
		have[s[i]-'a']++
		have[s[tail]-'a']--
		tail++
	}

	if have == need {
		res = append(res, tail)
	}

	return res
}
// func findAnagrams(s string, p string) []int {
// 	m := make(map[byte]int)
// 	tempM := make(map[byte]int)
// 	startArr := []int{}

// 	for i := 0; i < len(p); i++ {
// 		m[p[i]]++
// 		tempM[p[i]]++
// 	}

// 	left, right := 0, 0
// 	for right < len(s) {
// 		// check right
// 		if _, ok := tempM[s[right]]; !ok {
// 			left, right = right+1, right+1
// 			deepCopy(m, tempM)
// 		} else {
// 			tempM[s[right]]--

// 			if tempM[s[right]] < 0 {
// 				for ; left <= right; left++ {
// 					tempM[s[left]]++
// 					if s[left] == s[right] {
// 						left++
// 						break
// 					}
// 				}
// 			} else if right-left+1 == len(p) {
// 				startArr = append(startArr, left)
// 				tempM[s[left]]++
// 				left++
// 			}

// 			right++
// 		}

// 		fmt.Println(right, left, tempM)
// 	}

// 	return startArr
// }

// func deepCopy(m1, m2 map[byte]int) {
// 	for k, v := range m1 {
// 		m2[k] = v
// 	}
// }

// @lc code=end

