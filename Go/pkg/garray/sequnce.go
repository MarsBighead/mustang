package garray

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
// 121. 买卖股票的最佳时机
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	low, profit := 10000, 0
	for _, price := range prices {

		if price < low {
			low = price
		} else {
			tmp := price - low
			if tmp > profit {
				profit = tmp
			}
		}
	}
	return profit
}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
// 122. 买卖股票的最佳时机
// 充分利用了当天可以买卖的条件
func maxProfitv1(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += (prices[i] - prices[i-1])
		}
	}
	return profit
}
func maxProfitv2(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = 0 - prices[0]
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}
