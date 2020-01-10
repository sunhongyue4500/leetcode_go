package main

import "fmt"

func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
	fmt.Println(generateParenthesis(0))
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))

	fmt.Println(generateParenthesis_3(3))
	fmt.Println(generateParenthesis_4(3))
}

// 1. 使用暴力，对于n对，可以生成2n个位置，对于每一个位置，可以是(或者)
// 2. DFS
// 3. BFS
// 4. 动态规划
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

type Node struct {
	str         string
	leftUnUsed  int
	rightUnUsed int
}

// 3. BFS
func generateParenthesis_3(n int) []string {
	queue := []*Node{&Node{
		str:         "", // 当前组装的字符串
		leftUnUsed:  n,  //左括号剩余
		rightUnUsed: n,  //右括号剩余
	}}
	res := []string{}

	for len(queue) > 0 {
		node := queue[0]
		// pop
		queue = queue[1:]
		if node.leftUnUsed == 0 && node.rightUnUsed == 0 {
			res = append(res, node.str)
		}

		if node.leftUnUsed > 0 {
			queue = append(queue, &Node{
				str:         node.str + "(",
				leftUnUsed:  node.leftUnUsed - 1,
				rightUnUsed: node.rightUnUsed,
			})
		}
		// 右括号还有没用的并且左括号用的必须比右多
		if node.rightUnUsed > 0 && node.leftUnUsed < node.rightUnUsed {
			queue = append(queue, &Node{
				str:         node.str + ")",
				leftUnUsed:  node.leftUnUsed,
				rightUnUsed: node.rightUnUsed - 1,
			})
		}
	}
	return res
}

// 使用动态规划
// dp[i]，使用i对括号能生成的组合
// dp[i] = "(" + dp[可能的括号对数] + ")" + dp[剩下的括号对数]
// “可能的括号对数” 与 “剩下的括号对数” 之和得为 i - 1，故“可能的括号对数” j 可以从 0 开始，最大为i - 1；
// “剩下的括号对数” + j = i - 1，故 “剩下的括号对数” = i - j - 1。
func generateParenthesis_4(n int) []string {
	if n == 0 {
		return nil
	}
	dp := [][]string{}
	dp0 := []string{""}
	dp = append(dp, dp0)

	for i := 1; i <= n; i++ {
		cur := []string{}
		for j := 0; j < i; j++ {
			str1 := dp[j]
			str2 := dp[i-1-j]
			for _, s1 := range str1 {
				for _, s2 := range str2 {
					cur = append(cur, "("+s1+")"+s2)
				}
			}
		}
		dp = append(dp, cur)
	}
	return dp[n]
}
