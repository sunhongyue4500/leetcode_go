package main

import "fmt"

func main() {

	test1 := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	test2 := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	//fmt.Println(test1[0][0] == '5')
	//fmt.Printf("%s", test1)
	//fmt.Println(isValid2(test1, '5', 2, 0))

	//solveSudoku(test1)
	solveSudoku_2(test1)
	//solveSudoku_jianzhi(test2)
	solveSudoku_jianzhi_2(test2)
	fmt.Printf("%s", test1)
	fmt.Println()
	fmt.Printf("%s", test2)
}

// 这种写法不行不行
//func solveSudoku(board [][]byte) {
//	temp := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
//	dfs := func(board [][]byte, row, col int) bool { return false }
//	dfs = func(board [][]byte, row, col int) bool {
//		for i := row; i < 9; i++ {
//			for j := col; j < 9; j++ {
//				if board[i][j] == '.' {
//					// 尝试着填入
//					for k := 0; k < 9; k++ {
//						if isValid(board, temp[k], i, j) {
//							board[i][j] = temp[k]
//							if j >= 8 {
//								if dfs(board, i+1, 0) {
//									// 下一层
//									return true
//								} else {
//									board[i][j] = '.'
//								}
//							} else {
//								if dfs(board, i, j+1) {
//									// 下一层
//									return true
//								} else {
//									board[i][j] = '.'
//								}
//							}
//						}
//					}
//					return false
//				}
//			}
//		}
//		return true
//	}
//	fmt.Println(dfs(board, 0, 0))
//}

// 1. 朴素DFS，直接判断
func solveSudoku_2(board [][]byte) {
	temp := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	dfs := func(board [][]byte) bool { return false }
	dfs = func(board [][]byte) bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					// 尝试着填入
					for k := 0; k < 9; k++ {
						if isValid(board, temp[k], i, j) {
							board[i][j] = temp[k]
							if dfs(board) {
								// 下一层
								return true
							} else {
								board[i][j] = '.'
							}
						}
					}
					return false
				}
			}
		}
		return true
	}
	fmt.Println(dfs(board))
}

// target 填充到行为row，列为col的位置是否合适
func isValid(board [][]byte, target byte, row, col int) bool {
	// 检查所在的行
	for j := 0; j < 9; j++ {
		if j != col && board[row][j] != '.' && board[row][j] == target {
			return false
		}
	}
	// 检查所在的列
	for i := 0; i < 9; i++ {
		if i != row && board[i][col] != '.' && board[i][col] == target {
			return false
		}
	}
	// 检查所在的3*3方格
	tRow := row / 3
	tCol := col / 3
	for i := tRow * 3; i < (tRow+1)*3; i++ {
		for j := tCol * 3; j < (tCol+1)*3; j++ {
			if i != row && j != col && board[i][j] != '.' && board[i][j] == target {
				return false
			}
		}
	}
	return true
}

func isValid2(board [][]byte, target byte, row, col int) bool {
	for j := 0; j < 9; j++ {
		if board[row][j] != '.' && board[row][j] == target {
			return false
		}
		if board[j][col] != '.' && board[j][col] == target {
			return false
		}
		if board[3*(row/3)+j/3][3*(col/3)+j%3] != '.' && board[3*(row/3)+j/3][3*(col/3)+j%3] == target {
			return false
		}
	}
	return true
}

// DFS
// 2. 加入剪枝，用set这个数据结构，不允许重复
func solveSudoku_jianzhi(board [][]byte) {
	temp := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	// 索引代表行
	rowsSet := [9]map[byte]bool{}
	colsSet := [9]map[byte]bool{}
	blocksSet := [9]map[byte]bool{}

	// 初始化
	for i := 0; i < 9; i++ {
		rowsSet[i], colsSet[i], blocksSet[i] = make(map[byte]bool), make(map[byte]bool), make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				rowsSet[i][board[i][j]], colsSet[j][board[i][j]], blocksSet[i/3*3+j/3][board[i][j]] = true, true, true
			}
		}
	}
	dfs_jianzhi := func(board [][]byte, i, j int) bool { return true }
	dfs_jianzhi = func(board [][]byte, i, j int) bool {
		// terminator
		if i > 8 {
			return true
		}
		var target byte
		target = board[i][j]
		if target == '.' {
			for k := 0; k < 9; k++ {
				if !rowsSet[i][temp[k]] && !colsSet[j][temp[k]] && !blocksSet[i/3*3+j/3][temp[k]] {
					board[i][j] = temp[k]
					rowsSet[i][temp[k]], colsSet[j][temp[k]], blocksSet[i/3*3+j/3][temp[k]] = true, true, true
					if j == 8 {
						// 下一行
						if dfs_jianzhi(board, i+1, 0) {
							return true
						} else {
							board[i][j] = '.'
							rowsSet[i][temp[k]], colsSet[j][temp[k]], blocksSet[i/3*3+j/3][temp[k]] = false, false, false
						}
					} else {
						// 下一行
						if dfs_jianzhi(board, i, j+1) {
							return true
						} else {
							board[i][j] = '.'
							rowsSet[i][temp[k]], colsSet[j][temp[k]], blocksSet[i/3*3+j/3][temp[k]] = false, false, false
						}
					}
				}
			}
			return false
		} else {
			// 当前有数字
			if j == 8 {
				return dfs_jianzhi(board, i+1, 0)
			}
			return dfs_jianzhi(board, i, j+1)
		}
	}
	fmt.Println(dfs_jianzhi(board, 0, 0))
}

// 3. 加入剪枝，用set这个数据结构，不允许重复。使用两层for DFS
func solveSudoku_jianzhi_2(board [][]byte) {
	temp := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	// 索引代表行
	rowsSet := [9]map[byte]bool{}
	colsSet := [9]map[byte]bool{}
	blocksSet := [9]map[byte]bool{}

	// 初始化
	for i := 0; i < 9; i++ {
		rowsSet[i], colsSet[i], blocksSet[i] = make(map[byte]bool), make(map[byte]bool), make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				rowsSet[i][board[i][j]], colsSet[j][board[i][j]], blocksSet[i/3*3+j/3][board[i][j]] = true, true, true
			}
		}
	}

	dfs_jianzhi := func(board [][]byte) bool { return true }
	dfs_jianzhi = func(board [][]byte) bool {
		// terminator
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				var target byte
				target = board[i][j]
				if target == '.' {
					for k := 0; k < 9; k++ {
						if !rowsSet[i][temp[k]] && !colsSet[j][temp[k]] && !blocksSet[i/3*3+j/3][temp[k]] {
							board[i][j] = temp[k]
							rowsSet[i][temp[k]], colsSet[j][temp[k]], blocksSet[i/3*3+j/3][temp[k]] = true, true, true
							// 下一行
							if dfs_jianzhi(board) {
								return true
							} else {
								board[i][j] = '.'
								rowsSet[i][temp[k]], colsSet[j][temp[k]], blocksSet[i/3*3+j/3][temp[k]] = false, false, false
							}
						}
					}
					return false
				}
			}
		}
		return true
	}
	fmt.Println(dfs_jianzhi(board))
}
