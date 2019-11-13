package main

import "fmt"

func main() {
	//fmt.Println(totalNQueens(-6))
	//fmt.Println(totalNQueens(1))
	//fmt.Println(totalNQueens(2))
	//fmt.Println(totalNQueens(3))
	//fmt.Println(totalNQueens(4))
	//fmt.Println(totalNQueens(5))

	fmt.Println(totalNQueens_2(-6))
	fmt.Println(totalNQueens_2(1))
	fmt.Println(totalNQueens_2(2))
	fmt.Println(totalNQueens_2(3))
	fmt.Println(totalNQueens_2(4))
	fmt.Println(totalNQueens_2(5))
}

func totalNQueens(n int) int {
	if n == 1 {
		return 1
	}
	if n <= 3 {
		return 0
	}
	count := 0
	dfs := func(n, row int) {}
	col, pie, na := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	dfs = func(n, row int) {
		if row == n {
			count++
		}
		for j := 0; j < n; j++ {
			// 剪枝
			if !col[j] && !pie[row+j] && !na[row-j] {
				col[j], pie[row+j], na[row-j] = true, true, true
				dfs(n, row+1)
				col[j], pie[row+j], na[row-j] = false, false, false
			}
		}
	}
	dfs(n, 0)
	return count
}

func totalNQueens_2(n int) int {
	if n == 1 {
		return 1
	}
	if n <= 3 {
		return 0
	}
	col, pie, na := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	res := 0
	helper(n, 0, col, pie, na, &res)
	return res
}

func helper(n, row int, col, pie, na map[int]bool, count *int) {
	if row == n {
		*count++
	}
	for j := 0; j < n; j++ {
		// 剪枝
		if !col[j] && !pie[row+j] && !na[row-j] {
			col[j], pie[row+j], na[row-j] = true, true, true
			helper(n, row+1, col, pie, na, count)
			col[j], pie[row+j], na[row-j] = false, false, false
		}
	}
}
