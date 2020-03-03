package main

import "fmt"

func main() {
	fmt.Println(numDecodings(""))
	fmt.Println(numDecodings("1"))
	fmt.Println(numDecodings("17"))
	fmt.Println(numDecodings("012"))
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("123"))
	fmt.Println(numDecodings("1223"))

	fmt.Println()
	fmt.Println(numDecodings2("0"))
	fmt.Println(numDecodings2("1"))
	fmt.Println(numDecodings2("17"))
	fmt.Println(numDecodings("012"))
	fmt.Println(numDecodings2("123"))
	fmt.Println(numDecodings2("1223"))
}

// dp[i]表示包含几个字符
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1 // dp[0]为什么是1
	dp[1] = 1
	if s[0] == '0' {
		dp[1] = 0
	}
	for i := 2; i < len(s)+1; i++ {
		if s[i-1] > '0' && s[i-1] <= '9' {
			dp[i] = dp[i] + dp[i-1]
		}
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] < '7') {
			dp[i] = dp[i] + dp[i-2]
		}
	}
	return dp[len(s)]
}

func numDecodings2(s string) int {
	if len(s) == 0 {
		return 0
	}

	llast := 1
	last := 1
	if s[0] == '0' {
		last = 0
	}
	cur := last
	for i := 2; i < len(s)+1; i++ {
		if s[i-1] > '0' && s[i-1] <= '9' {
			cur = last
		}
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] < '7') {
			cur = cur + llast
		}
		llast = last
		last = cur
	}
	return cur
}
