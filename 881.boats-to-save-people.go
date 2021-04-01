/*
 * @lc app=leetcode id=881 lang=golang
 *
 * [881] Boats to Save People
 */

// @lc code=start
func numRescueBoats(people []int, limit int) int {
	sort.Sort(sort.IntSlice(people))

	i, j, count := 0, len(people)-1, 0

	for ; j > 0 && people[j] > limit; j-- {
	}

	for i < j {
		if people[i]+people[j] > limit {
			j--
		} else {
			i, j = i+1, j-1
		}
		count++
	}

	if i == j {
		count++
	}
	return count
}

// @lc code=end

