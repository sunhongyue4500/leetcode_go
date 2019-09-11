package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res1 [][]int

func main() {
	tree := prepareData()
	//fmt.Println(levelOrder(tree))

	//res := make(map[int][]int, 0)

	res1 = make([][]int, 0)
	dfs(tree, 0, &res1)

	fmt.Println(res1)

	//result := make([][]int, 0)
	//// map to array
	//for i:=0; i<len(res); i++ {
	//	result = append(result, res[i])
	//}
	//for key, val := range res {
	//	fmt.Println(key, val)
	//	result = append(result, val)
	//}

	//fmt.Println(levelOrder(tree))

}

func prepareData() *TreeNode {
	T1 := TreeNode{Val: 3}
	//var T1 TreeNode
	T2 := TreeNode{Val: 9}
	T3 := TreeNode{Val: 20}
	T4 := TreeNode{Val: 15}
	T5 := TreeNode{Val: 7}
	//
	T1.Left = &T2
	T1.Right = &T3
	T3.Left = &T4
	T3.Right = &T5
	return &T1
}

// BFS 不需要level
//func levelOrder(root *TreeNode) [][]int {
//	if root == nil {
//		return nil
//	}
//	queue := []*TreeNode{root}
//	res := make([][]int, 0)
//	var tempNode *TreeNode
//	for len(queue) != 0 {
//		// 保存这一层级
//		temp := make([]int, 0)
//		lenQueue := len(queue)
//		for i:=0; i<lenQueue; i++ { // 将这一层级全部出队
//			tempNode = queue[0]
//			queue = queue[1:] // pop
//			temp = append(temp, tempNode.Val)
//			if tempNode.Left != nil {
//				queue = append(queue, tempNode.Left)
//			}
//			if tempNode.Right != nil {
//				queue = append(queue, tempNode.Right)
//			}
//		}
//		// 将整个层级放到结果slice
//		res = append(res, temp)
//	}
//	return res
//}

// 使用切片时候要小心，append会重新分配
//func dfs(cu *TreeNode, level int, res map[int][]int) {
//	if cu == nil {
//		return
//	}
//	fmt.Println(cu.Val, level)
//	if _, ok := res[level]; !ok {
//		res[level] = make([]int, 0)
//	}
//	res[level] = append(res[level], cu.Val)
//
//	dfs(cu.Left, level + 1, res)
//	dfs(cu.Right, level + 1, res)
//}

// 全局变量版
//func dfs(cu *TreeNode, level int) {
//	if cu == nil {
//		return
//	}
//	fmt.Println(cu.Val, level)
//	if len(res1) - 1 < level {
//		res1 = append(res1, make([]int, 0))
//	}
//	res1[level] = append(res1[level], cu.Val)
//
//	dfs(cu.Left, level + 1)
//	dfs(cu.Right, level + 1)
//}

// DFS 切片指针
func dfs(cu *TreeNode, level int, res *[][]int) {
	if cu == nil {
		return
	}
	fmt.Println(cu.Val, level)
	if len(*res)-1 < level {
		*res = append(*res, make([]int, 0))
	}
	(*res)[level] = append((*res)[level], cu.Val)

	dfs(cu.Left, level+1, res)
	dfs(cu.Right, level+1, res)
}
