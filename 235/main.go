package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   15,
			Left:  &TreeNode{11, nil, nil},
			Right: &TreeNode{20, nil, nil},
		},
	}
	res := lowestCommonAncestor_1(&node, &TreeNode{
		Val:   20,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   11,
		Left:  nil,
		Right: nil,
	})
	fmt.Println(res.Val)
	fmt.Println(lowestCommonAncestor_1(nil, nil, nil))

	fmt.Println()
	res2 := lowestCommonAncestor_2(&TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   15,
			Left:  &TreeNode{11, nil, nil},
			Right: &TreeNode{20, nil, nil},
		},
	}, &TreeNode{
		Val:   20,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   11,
		Left:  nil,
		Right: nil,
	})
	fmt.Println(res2.Val)
	fmt.Println(lowestCommonAncestor_2(nil, nil, nil))
}

// 递归，利用BST性质
func lowestCommonAncestor_1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 找到了
	if (p.Val <= root.Val && q.Val >= root.Val) || (p.Val >= root.Val && q.Val <= root.Val) {
		return root
	}
	// 往左找
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor_1(root.Left, p, q)
	}
	return lowestCommonAncestor_1(root.Right, p, q)
}

// 代码结构优化
func lowestCommonAncestor_1_1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 往左找
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor_1_1(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor_1_1(root.Right, p, q)
	}
	return root
}

// 迭代，利用BST性质
func lowestCommonAncestor_2(root, p, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if (p.Val <= cur.Val && q.Val >= cur.Val) || (p.Val >= cur.Val && q.Val <= cur.Val) {
			return cur
		}
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return cur
}

// 迭代代码优化
func lowestCommonAncestor_2_2(root, p, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
		}
	}
	return cur
}
