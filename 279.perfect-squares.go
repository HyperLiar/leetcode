/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start
func numSquares(n int) int {
	// dp := make([]int, n+1)
	// for i := 1; i <= n; i++ {
	// 	dp[i] = dp[i-1] + 1
	// 	for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
	// 		//fmt.Println(maxSquare, maxProduct, i, j)
	// 		dp[i] = min(dp[i], dp[i-j*j]+1)
	// 	}
	// }

	// return dp[n]
	for n%4 == 0 {
		n /= 4
	}

	if n%8 == 7 {
		return 4
	}

	for i := 0; i*i <= n; i++ {
		j := int(math.Sqrt(float64(n - i*i)))
		if i*i+j*j == n {
			if i > 0 && j > 0 {
				return 2
			} else {
				return 1
			}
		}
	}

	return 3
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

// @lc code=end

