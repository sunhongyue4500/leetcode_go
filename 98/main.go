package main

import (
	"fmt"
	"math"
)

func main()  {
	//node1 := TreeNode{
	//	Val:2,
	//	Left:&TreeNode{1, nil, nil},
	//	Right:&TreeNode{3, nil, nil},
	//}
	//fmt.Println(isValidBST(&node1))	// true

	//node := TreeNode{
	//	Val:1,
	//	Left:&TreeNode{3, nil, nil},
	//	Right:&TreeNode{2, nil, nil},
	//}
	//fmt.Println(isValidBST(&node))	// false
	//fmt.Println(isValidBST(nil))
	//fmt.Println(isValidBST(&TreeNode{}))
	node := TreeNode{
		Val:10,
		Left:&TreeNode{
			Val:5,
			Left:nil,
			Right:nil,
		},
		Right:&TreeNode{
			Val:15,
			Left:&TreeNode{6, nil, nil},
			Right:&TreeNode{20, nil, nil},
		},
	}
	//fmt.Println(isValidBST_2(&node))
	fmt.Println(isValidBST_2(nil))
	fmt.Println(isValidBST_2(&TreeNode{}))
	fmt.Println(isValidBST_2(&TreeNode{0, nil, nil}))

	fmt.Println()
	fmt.Println(isValidBST_3(&node))
	fmt.Println(isValidBST_3(nil))
	fmt.Println(isValidBST_3(&TreeNode{}))
	fmt.Println(isValidBST_3(&TreeNode{0, nil, nil}))

	// 递归优化
	fmt.Println()
	fmt.Println(isValidBST_4(&node))
	fmt.Println(isValidBST_4(nil))
	fmt.Println(isValidBST_4(&TreeNode{}))
	fmt.Println(isValidBST_4(&TreeNode{0, nil, nil}))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var arr []int
func isValidBST(root *TreeNode) bool {
	// terminator
	arr = make([]int, 0)
	inorder(root)
	for i:=1; i<len(arr); i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	return true
}

func inorder(root *TreeNode) {
	// 左根右
	if root == nil {
		return
	}
	inorder(root.Left)
	arr = append(arr, root.Val)
	inorder(root.Right)
}

// 中序遍历优化，只保留前一个节点
var lastNode *TreeNode
func isValidBST_2(root *TreeNode) bool {
	lastNode = nil
	return helper(root)
}

func helper(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !helper(root.Left) {
		return false
	}
	// 当前节点
	if lastNode != nil && lastNode.Val >= root.Val {
		return false
	}
	lastNode = root // 记录
	return helper(root.Right) // 判断右节点
}

func isValidBST_3(root *TreeNode) bool {
	return helper2(root, math.MinInt64, math.MaxInt64)
}

// 解法2：直接递归，传左子树最大，和右子树最小
func helper2(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return helper2(root.Left, min, root.Val) && helper2(root.Right, root.Val, max)
}

// 解法2递归优化
func isValidBST_4(root *TreeNode) bool {
	return check(root, -1 << 63, 1 << 63 - 1)
}

func check(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	return root.Val > min && root.Val < max && check(root.Left, min, root.Val) && check(root.Right, root.Val, max)
}