/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	dp := make([][2]int, len(heights))
	for i := 0; i < len(heights); i++ {
		dp[i] = [2]int{i, i}
	}

	for i := 1; i < len(heights); i++ {
		left, right := i, i

        for dp[left][0] - 1 >= 0 && heights[dp[left][0] - 1] >= heights[i] {
			left = dp[left][0]-1
		}

		for right < len(heights) && heights[right] >= heights[i] {
			right++
		}
		dp[i][1] = right - 1

			
		for left >= 0 && heights[left] >= heights[i] {
			left--
		}
		dp[i][0] = left + 1
	}
	fmt.Println(dp)

	max := heights[0]
	for i := 0; i < len(heights); i++ {
		area := (dp[i][1] - dp[i][0] + 1) * heights[i]
		if area > max {
			max = area
		}
	}

	return max
}

// @lc code=end

