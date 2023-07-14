// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit := 0
	buy := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] > buy {
			maxProfit += prices[i] - buy
		}
		buy = prices[i]
	}
	return maxProfit
}
