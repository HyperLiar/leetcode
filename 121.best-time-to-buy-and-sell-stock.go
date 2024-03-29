/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	max, min := 0, 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[min] {
			min = i
		} else if cur := prices[i] - prices[min]; cur > max {
			max = cur
		}
	}

	return max
}
// @lc code=end

