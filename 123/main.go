package main

import (
	"log"
)

func main() {
	input := []int{
		1, 2, 4, 2, 5, 7, 2, 4, 9, 0,
	}

	log.Println(maxProfit(input))
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
