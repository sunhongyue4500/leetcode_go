package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "abc"))
	fmt.Println(longestCommonSubsequence("abc", "def"))
	fmt.Println(longestCommonSubsequence("h", "a"))

	fmt.Println(longestCommonSubsequence2("abcde", "ace"))
	fmt.Println(longestCommonSubsequence2("abc", "abc"))
	fmt.Println(longestCommonSubsequence2("abc", "def"))
	fmt.Println(longestCommonSubsequence2("h", "a"))

	fmt.Println()
	fmt.Println(longestCommonSubsequence3("abcde", "ace"))
	fmt.Println(longestCommonSubsequence3("abc", "abc"))
	fmt.Println(longestCommonSubsequence3("abc", "def"))
	fmt.Println(longestCommonSubsequence3("h", "a"))
}

// a dcka
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		dp[i] = make([]int, len(text2))
	}

	idx := strings.Index(text2, text1[0:1])
	if idx >= 0 {
		for i := idx; i < len(text2); i++ {
			dp[0][i] = 1
		}
	}
	idx = strings.Index(text1, text2[0:1])
	if idx >= 0 {
		for i := idx; i < len(text1); i++ {
			dp[i][0] = 1
		}
	}

	for i := 1; i < len(text1); i++ {
		for j := 1; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)-1][len(text2)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 优化，行列多加一层
func longestCommonSubsequence2(text1 string, text2 string) int {
	n1 := len(text1)
	n2 := len(text2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n1][n2]
}

// 状态压缩，只需要记录i-1，和i层，时间复杂度O(mn)，空间复杂度O(min(m, n))
func longestCommonSubsequence3(text1 string, text2 string) int {
	n1, n2 := len(text1), len(text2)
	if n1 > n2 {
		return longestCommonSubsequence3(text2, text1)
	}
	dp := [2][]int{make([]int, n2+1), make([]int, n2+1)}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i%2][j] = dp[(i-1)%2][j-1] + 1
			} else {
				dp[i%2][j] = max(dp[(i-1)%2][j], dp[i%2][j-1])
			}
		}
	}
	return dp[n1%2][n2]
}
