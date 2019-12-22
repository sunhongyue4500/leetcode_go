package main

import "fmt"

func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
	fmt.Println(generateParenthesis(0))
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
}

// 1. 使用暴力，对于n对，可以生成2n个位置，对于每一个位置，可以是(或者)
// 2. 递归
func generateParenthesis(n int) []string {
	res := []string{}
	helper(n, 0, 0, "", &res)
	return res
}

// 该递归函数表示，当括号对数为n的时候，已经使用了左括号leftUsed，使用了右括号rightUsed个，
// 当前已拼接的字符串状态为str，结果集为res
func helper(n int, leftUsed int, rightUsed int, str string, res *[]string) {
	if leftUsed == n && rightUsed == n {
		// 都用光了
		*res = append(*res, str)
		return
	}
	if leftUsed < n {
		helper(n, leftUsed+1, rightUsed, str+"(", res)
	}
	if leftUsed > rightUsed && rightUsed < n {
		helper(n, leftUsed, rightUsed+1, str+")", res)
	}
}

func generateParenthesis2(n int) []string {
	res := []string{}
	helper(n, n, n, "", &res)
	return res
}

func helper2(n int, left int, right int, curStr string, res *[]string) {
	// terminator
	if left > right {
		return
	}
	// process data
	if left == 0 && right == 0 {
		*res = append(*res, curStr)
		return
	}

	// drill down
	if left >= 1 {
		helper(n, left-1, right, curStr+"(", res)
	}
	if right >= 1 {
		helper(n, left, right-1, curStr+")", res)
	}
	// clear status
}
