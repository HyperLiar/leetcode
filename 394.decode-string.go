/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func decodeString(s string) string {
	if len(s) <= 1 {
		return s
	}
	res := ""

	var prefix, next string
	var num, left, right, idx int
	for ; idx < len(s); idx++ {
		if int(s[idx]-'a') > 26 || int(s[idx]-'a') < 0 {
			break
		}
	}
	if idx >= len(s)-1 {
		return s
	}
	//fmt.Println("prefix", s, prefix, idx)
	prefix = s[0:idx]

	left = idx
	for ; left < len(s); left++ {
		if s[left] == '[' {
			break
		}
	}
	//fmt.Println("left", s, left, idx)


	if left == len(s)-1 {
		num = 1
		next = s[idx:]
		right = len(s) - 2
	} else {
		//fmt.Println("num", s[idx:left])

		num, _ = strconv.Atoi(s[idx:left])
		leftCount, rightCount := 1, 0
		right = left + 1
		for ; right < len(s); right++ {
			if s[right] == '[' {
				leftCount++
			} else if s[right] == ']' {
				rightCount++
			}

			if leftCount == rightCount {
				break
			}
		}
		//fmt.Println("next", left, right, s[left+1 : right])

		next = decodeString(s[left+1 : right])
	}

	//fmt.Println(idx, num, left, right, s)

	res += prefix
	for i := 0; i < num; i++ {
		res += next
	}
	res += decodeString(s[right+1:])

	return res
}

// @lc code=end

