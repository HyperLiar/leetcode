/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {
	n, res := len(height), 0
	l, r, lmax, rmax := 0, n-1,0, 0

	for l < r {
		if height[l] < height[r] {
			if height[l] >= lmax {
				lmax = height[l]
			} else {
				res += lmax - height[l]
			}
			l++
		} else {
			if height[r] >= rmax {
				rmax = height[r]
			} else {
				res += rmax - height[r]
			}
			r--
		}
	}

	return res
}

func temp(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	maxLeft, maxRight, res := height[0], height[n-1], 0
	// left
	left, right := 1, n-2
	for ; left < n; left++ {
		if height[left] >= maxLeft {
			maxLeft = height[left]
		} else {
			break
		}
	}
	left--
	for ; right >= 0; right-- {
		if height[right] >= maxRight {
			maxRight = height[right]
		} else {
			break
		}
	}
	right++
	if right-left < 2 {
		return 0
	}

	i, j := left+1, right-1
	for i <= j {
		if height[i] > maxLeft {
			maxLeft = height[i]
			left = i
		} else {
			res += maxLeft - height[i]
		}
		if height[j] > maxRight {
			maxRight = height[j]
			right = j
		} else {
			res += maxRight - height[j]
		}

		i++
		j--
	}

	//fmt.Println(maxLeft, maxRight, right, left, i, j, res)
	if maxLeft == maxRight {
		if i == j+2 {
			res -= maxLeft - height[i-1]
		}
		return res
	}

	for x := left + 1; x < i; x++ {
		res -= maxLeft - height[x]
	}
	for x := right - 1; x > j; x-- {
		res -= maxRight - height[x]
	}

	if maxLeft < maxRight {
		for j := left + 1; j < right; j++ {
			if height[j] > maxLeft {
				maxLeft = height[j]
			} else {
				res += maxLeft - height[j]
			}
		}
	} else {
		for j := right - 1; j > left; j-- {
			if height[j] > maxRight {
				maxRight = height[j]
			} else {
				res += maxRight - height[j]
			}
		}
	}

	return res
}
// @lc code=end

