package main

import "fmt"

func main() {

	test1 := [][]byte{
		[]byte{'1', '3', '2', '.', '7', '.', '.', '.', '.'},
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

	fmt.Println(isValidSudoku(test1))
}

func isValidSudoku(board [][]byte) bool {
	var target byte
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			target = board[i][j]
			if target != '.' {
				// 判断是否有效
				for k := 0; k < 9; k++ {
					// 检测行
					if k != j && board[i][k] != '.' && board[i][k] == target {
						return false
					}
					// 检测列
					if i != k && board[k][j] != '.' && board[k][j] == target {
						return false
					}
				}
				// 测试3*3方块
				for m := i / 3 * 3; m < (i/3+1)*3; m++ {
					for n := j / 3 * 3; n < (j/3+1)*3; n++ {
						if m != i && n != j && board[m][n] != '.' && board[m][n] == target {
							return false
						}
					}
				}
			}
		}
	}
	return true
}
