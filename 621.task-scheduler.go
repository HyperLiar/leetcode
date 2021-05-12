/*
 * @lc app=leetcode id=621 lang=golang
 *
 * [621] Task Scheduler
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	s := [26]int{}
	for i := 0; i < len(tasks); i++ {
		s[int(tasks[i]-'A')]++
	}

	max_count := 0

	for i := 0; i < len(s); i++ {
		if s[i] > max_count {
			max_count = s[i]
		}
	}

	res := (max_count - 1) * (n+1)

	for i := 0; i < len(s); i++ {
		if s[i] == max_count {
			res++
		}
	}

	if len(tasks) > res {
		return len(tasks)
	}
	return res
}

// @lc code=end

