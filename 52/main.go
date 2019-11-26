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
	fmt.Println()
	fmt.Println(totalNQueens3(-6))
	fmt.Println(totalNQueens3(1))
	fmt.Println(totalNQueens3(2))
	fmt.Println(totalNQueens3(3))
	fmt.Println(totalNQueens3(4))
	fmt.Println(totalNQueens3(5))
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

/********************使用位运算*********************/
func totalNQueens3(n int) int {
	if n == 1 {
		return 1
	}
	if n <= 3 {
		return 0
	}
	cnt := 0
	DFS(n, 0, 0, 0, 0, &cnt)
	return cnt
}

func DFS(n, row, col, pie, na int, count *int) {
	// terminator
	if row >= n {
		*count = *count + 1
		return
	}

	// ^在go中作为一元运算符是取反操作，a^b则是异或
	// 列不是从头扫到尾，而是通过位运算快速得到空位
	// (col | pie | na)表示将列撇捺汇总，空位即是0
	// 上述取反，就是找到1为空位，这时候还包括高位的1
	// &（(1<<n）-1) 表示将二进制高位的1去掉，如果是4皇后就表示 & (0000 1111)，起到一个筛子的作用
	bits := (^(col | pie | na)) & ((1 << n) - 1)
	// 尝试可能放的位置，枚举所有的二进制位为1
	for bits != 0 {
		// 取到最低位的1，即表示可以放queen的位置
		p := bits & -bits
		// col | p，表示col放上p这一位，
		// 进入到下一行，pie要左移一位，捺要右移一位
		DFS(n, row+1, col|p, (pie|p)<<1, (na|p)>>1, count)
		bits = bits & (bits - 1)
	}
}
