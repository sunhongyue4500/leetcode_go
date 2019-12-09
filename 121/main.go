package main

import (
	"fmt"
	"math"
)

func main() {
	test := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(test))
	fmt.Println(maxProfit2(test))
	fmt.Println(maxProfit3(test))
}

// 1. 暴力，N^2
// 2. 找到最小的谷之后，找最大的峰

// [7,1,5,3,6,4]
func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

// 国际站的一种做法
func maxProfit3(prices []int) int {
	maxCur, maxSofar := 0, 0
	for i := 1; i < len(prices); i++ {
		maxCur = max(0, maxCur+prices[i]-prices[i-1])
		if maxCur > maxSofar {
			maxSofar = maxCur
		}
	}
	return maxSofar
}

// 用DP来做
func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	k := 1
	buy := make([]int, k+1)
	sell := make([]int, k+1)
	// 初始化
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
