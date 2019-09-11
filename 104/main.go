package main

import "fmt"

func main() {
	//fmt.Println(maxDepth(&TreeNode{1, &TreeNode{2, nil, &TreeNode{2, nil, nil}}, nil}))
	fmt.Println(maxDepth(nil))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
//func maxDepth(root *TreeNode) int {
//	res := 0
//	dfs(root, 0, &res)
//	fmt.Println(res)
//	return res
//}
//
//func dfs(root *TreeNode, level int, res *int) {
//	if root == nil {
//		return
//	}
//	level++
//	if *res < level {
//		*res = level
//	}
//	dfs(root.Left, level, res)
//	dfs(root.Right, level, res)
//}

// BFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	level, lenTemp := 0, 0
	var tempNode *TreeNode
	for len(queue) != 0 {
		lenTemp = len(queue)
		// 处理当前层
		for i := 0; i < lenTemp; i++ {
			// 将当前层全部pop
			tempNode = queue[0]
			queue = queue[1:]
			if tempNode.Left != nil {
				queue = append(queue, tempNode.Left)
			}
			if tempNode.Right != nil {
				queue = append(queue, tempNode.Right)
			}
		}
		level++
	}
	return level
}
