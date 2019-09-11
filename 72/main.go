package main

import (
	"fmt"
)

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	m := make([][]int, n1+1)

	for i := 0; i <= n1; i++ {
		m[i] = make([]int, n2+1)
		m[i][0] = i
	}
	for j := 0; j <= n2; j++ {
		m[0][j] = j
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				m[i][j] = m[i-1][j-1]
			} else {
				m[i][j] = getMin(m[i-1][j-1], m[i-1][j], m[i][j-1]) + 1
			}
		}
	}
	return m[n1][n2]
}

func getMin(a int, b int, c int) int {
	if a > b {
		if b > c {
			return c
		} else {
			return b
		}
	} else {
		if a < c {
			return a
		} else {
			return c
		}
	}
}
