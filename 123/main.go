package main

import (
	"fmt"
	"log"
)

func main() {
	input := []int{
		1, 2, 4, 2, 5, 7, 2, 4, 9, 0,
	}

	log.Println(maxProfit(input))
	fmt.Println(maxProfit2(input))
	fmt.Println(maxProfit3(input))
}

func maxProfit(prices []int) int {
	start, end, max, secmax := 0, 0, 0, 0
	for i := 1; i < len(prices); i++ {
		// 找到第一个最低点
		for j := i; j < len(prices); j++ {
			if prices[j] < prices[j-1] {
				start = j
			} else {
				break
			}
		}
		end = start + 1
		for j := end; j < len(prices); j++ {
			if prices[j] >= prices[j-1] {
				end = j
			} else {
				break
			}
		}
		if end < len(prices) {
			max, secmax = temp(prices[end]-prices[start], max, secmax)
		}

		i = end
	}
	return max + secmax
}

func temp(cur int, max int, secmax int) (int, int) {
	if cur > max {
		return cur, max
	} else if cur > secmax {
		return max, cur
	}
	return max, secmax
}

// DP
func maxProfit2(prices []int) int {
	// boundary conditions
	if len(prices) < 2 {
		return 0
	}
	k := 2
	buy := make([]int, k+1)  // 买入k次的最大利润
	sell := make([]int, k+1) // 卖出k次的最大利润
	// dp初始化
	for i := 0; i <= k; i++ {
		buy[i] = -(1 << 63)
	}
	for i := 0; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}
	return sell[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 优化的解法
func maxProfit3(prices []int) int {
	oneBuy, oneBuyOneSell, twoBuy, twoBuyTwoSell := -(1 << 63), 0, -(1 << 63), 0
	for i := 0; i < len(prices); i++ {
		oneBuy = max(oneBuy, -prices[i])
		oneBuyOneSell = max(oneBuyOneSell, oneBuy+prices[i])
		twoBuy = max(twoBuy, oneBuyOneSell-prices[i])
		twoBuyTwoSell = max(twoBuyTwoSell, twoBuy+prices[i])
	}
	return twoBuyTwoSell
}
