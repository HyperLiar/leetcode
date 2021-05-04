/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	l := len(heights)
	if l == 0 {
		return 0
	}
	
	dpLeft := make([]int, l)
	dpRight := make([]int, l)

	for i := 1; i < l; i++ {
		dpLeft[i] = i
		j := i

		for dpLeft[j] - 1 >= 0 && heights[dpLeft[j] - 1] >= heights[i] {
            j = dpLeft[j] - 1
        }
        dpLeft[i] = dpLeft[j]
	}

	dpRight[l-1] = l-1

	for i := l-2; i >= 0; i-- {
		dpRight[i] = i
		j := i

		for dpRight[j] + 1 <= l-1 && heights[dpRight[j]+1] >= heights[i] {
			j = dpRight[j]+1
		}
		dpRight[i] = dpRight[j]
	}

	//fmt.Println(dpLeft, dpRight)

	result := 0
	for i := 0; i < l; i++ {
		x := heights[i] * (dpRight[i] - dpLeft[i] + 1) 
		if x > result {
			result = x
		}
	}

	return result
}

// @lc code=end

