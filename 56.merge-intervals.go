/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	// do sort
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1] {
			return true
		} else {
			return false
		}
	})

	res := make([][]int, 0)
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] < intervals[i+1][0] {
			res = append(res, intervals[i])
		} else {
			intervals[i+1][0] = intervals[i][0]
			if intervals[i][1] > intervals[i+1][1] {
				intervals[i+1][1] = intervals[i][1]
			}
		}
	}

	res = append(res, intervals[len(intervals)-1])

	return res
}

// @lc code=end

