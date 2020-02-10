package main

import "fmt"

func main() {
	fmt.Println(minCostClimbingStairs([]int{15, 20, 25}))
}

// 1. dp
// dp[i]表示走到索引为i的台阶的最低话费
// dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
// res = min(dp[i-1], dp[i-2])
// 空间压缩，仅使用变量
// 时间复杂度O(N)，空间复杂度O(1)
func minCostClimbingStairs(cost []int) int {
	last, cur := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		last, cur = cur, min(cur, last)+cost[i]
	}
	return min(cur, last)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
