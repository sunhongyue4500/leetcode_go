package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(test))
	fmt.Println(maxProfit2(test))
	fmt.Println(maxProfit3(test))
}

// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// dp[i][0, 1][0, 1] 表示第i天不持有和持有，是否是冷冻期的最大利润
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0], dp[0][1] = 0, 0-prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
		if res < dp[i][0] {
			res = dp[i][0]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// hold：表示继续持有股票
// sold：表示卖出股票
// reset：表示什么都不做
func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	buy := make([]int, len(prices))
	sell := make([]int, len(prices))
	cooldown := make([]int, len(prices))
	buy[0] = -prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		// 继续持有或者冷冻期后买入
		buy[i] = max(buy[i-1], cooldown[i-1]-prices[i])
		// 卖出股票
		sell[i] = buy[i-1] + prices[i]
		// 什么都不做:继续什么都不做，或者昨天卖出后什么都不做
		cooldown[i] = max(cooldown[i-1], sell[i-1])
		if res < sell[i] {
			res = sell[i]
		}
		if res < cooldown[i] {
			res = cooldown[i]
		}
	}
	return res
}

// DP 状态压缩
func maxProfit3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	hold, sold, reset, res, temp := -prices[0], 0, 0, 0, 0
	for i := 1; i < len(prices); i++ {
		temp = reset
		// 什么都不做:继续什么都不做，或者昨天卖出后什么都不做
		reset = max(reset, sold)
		// 卖出股票
		sold = hold + prices[i]
		// 继续持有或者冷冻期后买入
		hold = max(hold, temp-prices[i])
		temp = max(sold, reset)
		if res < temp {
			res = temp
		}
	}
	return res
}
