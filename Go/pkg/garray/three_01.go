package garray

import (
	"strconv"
	"strings"
)

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

// 134. 加油站
// https://leetcode.cn/problems/gas-station/description/?envType=study-plan-v2&envId=top-interview-150
func canCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}
	return -1

}

// 165. 比较版本号
// https://leetcode.cn/problems/compare-version-numbers/description/
func compareVersion(version1 string, version2 string) int {
	var (
		vs1 = strings.Split(version1, ".")
		vs2 = strings.Split(version2, ".")
	)
	m, n := len(vs1), len(vs2)
	length := n
	if m > n {
		length = m
	}
	var (
		a, b, i int
	)
	for ; i < length; i++ {
		if i < m {
			a, _ = strconv.Atoi(vs1[i])
		} else {
			a = 0
		}
		if i < n {
			b, _ = strconv.Atoi(vs2[i])
		} else {
			b = 0
		}
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
	}
	return 0
}

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
// 167. 两数之和 II - 输入有序数组
// 双指针逼近
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	if (target > 0 && numbers[i] > target) || (target < 0 && target > numbers[j]) {
		return []int{}
	}
	for i < j {
		v := numbers[i] + numbers[j]
		if v == target {
			return []int{i + 1, j + 1}
		} else if v < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}
