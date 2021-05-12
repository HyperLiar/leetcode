/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 */

// @lc code=start
type ss struct {
	start, end int
}

func partitionLabels(s string) []int {
	area, temp := [26]ss{}, make([]ss, 0)
	for i := 0; i < 26; i++ {
		area[i].start = -1
	}
	for i := 0; i < len(s); i++ {
		if area[s[i]-'a'].start == -1 {
			area[s[i]-'a'].start = i
		}

		if area[s[i]-'a'].end < i {
			area[s[i]-'a'].end = i
		}
	}

	for i := 0; i < 26; i++ {
		if area[i].start != -1 {
			temp = append(temp, area[i])
		}
	}

	sort.Slice(temp, func(i, j int) bool {
		if temp[i].start < temp[j].start {
			return true
		} else if temp[i].start == temp[j].start {
			return temp[i].end < temp[j].end
		}
		return false
	})

	res := []int{}
	for i := 0; i < len(temp)-1; i++ {
		if temp[i].end < temp[i+1].start {
			res = append(res, temp[i].end - temp[i].start + 1)
		} else if temp[i].end >= temp[i+1].end {
			temp[i+1].start = temp[i].start
			temp[i+1].end = temp[i].end
		} else {
			temp[i+1].start = temp[i].start
		}
	}

	return append(res, temp[len(temp)-1].end - temp[len(temp)-1].start + 1)
}

// @lc code=end

