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

// https://leetcode.cn/problems/jump-game/description
// 55. 跳跃游戏
func canJump(nums []int) bool {
	n, right := len(nums), 0
	if nums[0] == 0 && n > 1 {
		return false
	}
	/*
		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		for i, num := range nums {
			// i=0，计算right；若nums[0]=0,肯定false
			if i <= right {
				right = max(right, i+num)
				if right >= n-1 {
					return true
				}
			}
		}
		return false
	*/

	for i, num := range nums {
		if i <= right {
			if right < i+num {
				right = i + num
			}
		}
	}
	return right >= n-1
}

// https://leetcode.cn/problems/jump-game-ii/description
// 45. 跳跃游戏 II
// 正向查找可到达的最大位置，不计算最后一步
func jump(nums []int) int {
	var (
		pos   = 0
		steps = 0
		right = 0
	)
	for i := 0; i < len(nums)-1; i++ {
		if pos < i+nums[i] {
			pos = i + nums[i]
		}
		if i == right {
			right = pos
			steps++
		}
	}
	return steps
}

// https://leetcode.cn/problems/h-index/description
// 274. H 指数
func hIndex(citations []int) int {
	n := len(citations)
	hs := make([]int, n+1)
	for _, citation := range citations {
		if citation > n {
			citation = n
		}
		hs[citation] += 1
	}
	for i, total := n, 0; i >= 0; i-- {
		total += hs[i]
		if total >= i {
			return i
		}
	}

	return 0

}
