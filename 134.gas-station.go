/*
 * @lc app=leetcode id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		res, cnt := 0, 0
		for cnt < n {
			j := (i + cnt) % n
			res += gas[j] - cost[j]
			if res < 0 {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}

	return -1
}

// @lc code=end

