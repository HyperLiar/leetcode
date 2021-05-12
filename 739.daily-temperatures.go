/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// @lc code=start
type VAndI struct {
	v, idx int
}

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := []int{}
	n := len(temperatures)

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[0 : len(stack)-1]
		}

		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)
		//fmt.Println(stack)
	}

	return ans
}

// @lc code=end

