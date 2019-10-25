package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{9, nil, nil},
			Right: &TreeNode{2, nil, nil},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{0, nil, nil},
			Right: &TreeNode{8, nil, nil},
		},
	}

	fmt.Println("pre order")
	preOrder(tree)
	fmt.Println("in order")
	inOrder(tree)
	fmt.Println("post order")
	postOrder(tree)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}

// 如果是BST，中序遍历是有序的

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Println(root.Val)
	inOrder(root.Right)
}

func postOrder(root *TreeNode) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Println(root.Val)
}
