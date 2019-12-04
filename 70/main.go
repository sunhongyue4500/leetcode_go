package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs2(2))
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 状态压缩
func climbStairs2(n int) int {
	lastlast, last, res := 0, 1, 1
	for i := 1; i <= n; i++ {
		res = last + lastlast
		lastlast, last = last, res
	}
	return res
}
