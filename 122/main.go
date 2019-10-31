package main

import "fmt"

func main() {
	testData := []int{7, 1, 5, 2, 1, 9, 3, 6, 4}
	test2 := []int{1, 2, 3, 4}
	test3 := []int{4, 3, 2, 1}
	test4 := []int{6, 1, 3, 2, 4, 7}
	test5 := []int{1, 4, 2}
	test6 := []int{1, 2, 2, 4, 5, 5, 1, 1, 7}
	fmt.Println(maxProfit(testData))
	fmt.Println(maxProfit_2(testData))
	fmt.Println(maxProfit_2(test2))
	fmt.Println(maxProfit_2(test3))
	fmt.Println(maxProfit_2(test4))
	fmt.Println(maxProfit_2(test5))
	fmt.Println()
	fmt.Println(maxProfit_3(test2))
	fmt.Println(maxProfit_3(test3))
	fmt.Println(maxProfit_3(test4))
	fmt.Println(maxProfit_3(test5))
	fmt.Println(maxProfit_3(test6))
}

// 凡是有增涨就买
func maxProfit(prices []int) int {

	res := 0 // 利润

	for i := 1; i < len(prices); i++ {
		// 找到第一个下降点
		if prices[i]-prices[i-1] > 0 {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}

// 找波峰和波谷，这个代码写的比较丑
func maxProfit_2(prices []int) int {
	res := 0 // 利润
	low := 0
	high := -1
	for i := 0; i < len(prices); i++ {
		// 找到第一个下降点
		if i+1 < len(prices) && prices[i+1] < prices[i] {
			low = i + 1
		}
		// 要先找到最低点才行
		if i+1 < len(prices) && prices[i+1] > prices[i] {
			high = i + 1
		}
		if low >= 0 && high >= 0 && high > low {
			res += prices[high] - prices[low]
			high, low = -1, high // 重置
		}
	}
	if low < 0 && high >= 0 {
		return prices[high] - prices[0]
	}

	return res
}

// 优化，找波峰和波谷
func maxProfit_3(prices []int) int {
	vally, peak, res, i := 0, 0, 0, 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		vally = i

		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = i

		res += prices[peak] - prices[vally]
	}
	return res
}
