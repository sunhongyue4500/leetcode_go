package main

import "fmt"

func main() {
	ob := [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 1, 0, 0, 0, 1, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0},
		[]int{1, 0, 1, 0, 0, 1, 0, 0},
		[]int{0, 0, 1, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 1, 1, 0, 1, 0},
		[]int{0, 1, 0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(ob))
	fmt.Println(uniquePathsWithObstacles([][]int{[]int{0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{[]int{1}}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 定义状态 dp[i][j]表示，从start到i，j的位置有多少种路径
	// 状态转移方程：dp[i][j] = dp[i-1][j] + dp[i][j-1]，如果是石头，dp[i][j] = 0
	// 压缩状态变量
	if obstacleGrid == nil {
		return 0
	}
	row := len(obstacleGrid)
	col := len(obstacleGrid[0])

	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	// 初始化
	for i := 0; i < row; i++ {
		// 初始化的时候也要排除
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 0; i < col; i++ {
		// 初始化的时候也要排除
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[row-1][col-1]
}
