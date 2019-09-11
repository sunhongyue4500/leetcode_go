package main

import (
	"fmt"
)

func main() {
	input := [][]int{
		{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(input))
}

func minimumTotal2(triangle [][]int) int {
	// DFS
	res := DFS(0, 0, triangle, 0)
	fmt.Printf("res:%d", res)
	return 0
}

func DFS(i, j int, triangle [][]int, res int) int {
	// terminator
	if i == len(triangle)-1 {
		// 到达边界
		fmt.Printf("->%d \n", res+triangle[i][j])
		return res + triangle[i][j]
	}
	// process i, j
	sum := res + triangle[i][j]
	//fmt.Printf("进入(%d, %d) ", i, j)

	// drill down
	DFS(i+1, j, triangle, sum)
	DFS(i+1, j+1, triangle, sum)
	// clear status(no need)

	//fmt.Printf("最小：%d \n", sum)

	return sum
}

//func minimumTotal(triangle [][]int) int {
//	temp := triangle[len(triangle) - 1]
//	maptemp := make(map[string]int)
//	return recur(0, 0, temp, triangle, maptemp)
//}
//
//func recur(i int, j int, arr []int, tri [][]int, m map[string]int) int {
//	if v, ok := m[strconv.Itoa(i) + "-" + strconv.Itoa(j)]; ok {
//		return v
//	}
//	if i == len(tri) -1 {
//		m[strconv.Itoa(i) + "-" + strconv.Itoa(j)] = arr[j]
//		return arr[j]
//	}
//	temp1 := recur(i+1, j, arr, tri, m)
//	m[strconv.Itoa(i+1) + "-" + strconv.Itoa(j)] = temp1
//	temp2 := recur(i+1, j+1, arr, tri, m)
//	m[strconv.Itoa(i+1) + "-" + strconv.Itoa(j+1)] = temp2
//	return min(temp1, temp2) + tri[i][j]
//}
//
//func min(a int, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	temp := triangle[len(triangle)-1]
	return recur(0, 0, temp, triangle)
}

func recur(i int, j int, arr []int, tri [][]int) int {
	if i == len(tri)-1 {
		return arr[j]
	}
	// process
	temp1 := recur(i+1, j, arr, tri)
	temp2 := recur(i+1, j+1, arr, tri)
	res := min(temp1, temp2) + tri[i][j]
	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
