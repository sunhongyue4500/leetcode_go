package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs2(2))
	fmt.Println(climbStairs(10))
	fmt.Println(climbStairs2(10))
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 状态压缩，仅使用两个变量
func climbStairs2(n int) int {
	if n < 3 {
		return n
	}
	first, sec := 1, 2
	for i := 3; i <= n; i++ {
		first, sec = sec, first+sec
	}
	return sec
}
