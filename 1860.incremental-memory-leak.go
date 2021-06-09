/*
 * @lc app=leetcode id=1860 lang=golang
 *
 * [1860] Incremental Memory Leak
 */

// @lc code=start
func memLeak(memory1 int, memory2 int) []int {
	i := 1
	for memory1 >= 0 && memory2 >= 0 {
		if memory1 >= memory2 {
			memory1 -= i
		} else {
			memory2 -= i
		}
		i++
	}
	i--
	if memory1 < 0 {
		memory1 += i
	}
	if memory2 < 0 {
		memory2 += i
	}
	return []int{i, memory1, memory2}
}

// @lc code=end

