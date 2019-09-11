package main

import "fmt"

func main() {
	input := []int{1}
	input2 := 0
	fmt.Println(coinChange(input, input2))
}

func coinChange(coins []int, amount int) int {
	if amount <= 0 || len(coins) == 0 {
		return 0
	}
	max := 2 << 31
	dp := make([]int, amount+1)
	var temp int
	for i := 1; i < amount+1; i++ {
		temp = max
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] && temp > dp[i-coins[j]] && dp[i-coins[j]] >= 0 {
				// dp[i] = dp[i - coins[j]] + 1
				temp = dp[i-coins[j]]
			}
		}
		if temp < max {
			dp[i] = temp + 1
		} else {
			dp[i] = -1
		}
	}
	return dp[amount]
}
