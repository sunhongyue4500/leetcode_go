package main

import "fmt"

func main() {
	//fmt.Println(minDepth(&TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, &TreeNode{5, nil, nil}, nil}}, nil}))
	//fmt.Println(minDepth(nil))
	fmt.Println(minDepth(&TreeNode{2, &TreeNode{}, nil}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
//func minDepth(root *TreeNode) int {
//	//lc, lr := 0, 0
//	//if root == nil {
//	//	return 0
//	//}
//	//lc = maxDepth(root.Left)
//	//lr = maxDepth(root.Right)
//	//
//	//if lc < lr {
//	//	return lc + 1
//	//}
//	//return lr + 1
//
//	if root == nil {
//		return 0
//	}
//	res := 2 << 32
//	dfs(root, 0, &res)
//	return res
//}
//
//func dfs(root *TreeNode, level int, res *int) {
//	if root == nil || level+1 > *res {
//		return
//	}
//	level++
//	if root.Left == nil && root.Right == nil {
//		*res = level
//		return
//	}
//	dfs(root.Left, level, res)
//	dfs(root.Right, level, res)
//}

// BFS
func minDepth(root *TreeNode) int {
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
			if tempNode.Left == nil && tempNode.Right == nil {
				return level + 1
			}
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
