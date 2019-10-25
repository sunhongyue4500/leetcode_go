package main

import "fmt"

func main() {
	res2 := lowestCommonAncestor(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{0, nil, nil},
			Right: &TreeNode{8, nil, nil},
		},
	}, &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	})
	fmt.Println(res2.Val)
	fmt.Println(lowestCommonAncestor(nil, nil, nil))

	arr := make([]*TreeNode, 0)
	DFS(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{0, nil, nil},
			Right: &TreeNode{8, nil, nil},
		},
	}, &TreeNode{0, nil, nil}, &arr)

	test2()

	fmt.Println("lowestCommonAncestor_2")
	lowestCommonAncestor2Res := lowestCommonAncestor_2(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{0, nil, nil},
			Right: &TreeNode{8, nil, nil},
		},
	}, &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	})
	if lowestCommonAncestor2Res != nil {
		fmt.Println(lowestCommonAncestor2Res.Val)
	}

	fmt.Println("lowestCommonAncestor_3")
	//
	res3 := lowestCommonAncestor_3(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{0, nil, nil},
			Right: &TreeNode{8, nil, nil},
		},
	}, &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}, &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	})
	if res3 != nil {
		fmt.Println("res3:", res3.Val)
	}
	lowestCommonAncestor_3(nil, nil, nil)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return findPorQ(root, p, q)
}

// 表示在以root为根的树中找p，q，如果两个都找了，随便返回一个
// if root == p or root == q, 直接返回root
// 左子树去找，如果都在左子树，就不去右子树了
func findPorQ(root, p, q *TreeNode) *TreeNode {
	// terminator
	if root == nil || p == nil || q == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// Drill Down
	left := findPorQ(root.Left, p, q)
	right := findPorQ(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度优先搜索
func DFS(root *TreeNode, p *TreeNode, arr *[]*TreeNode) {
	// terminator
	if root == nil {
		fmt.Println()
		return
	}
	if root.Val == p.Val {
		*arr = append(*arr, root)
	}
	// process
	fmt.Print(root.Val, "-")
	DFS(root.Left, p, arr)
	DFS(root.Right, p, arr)
}

func lowestCommonAncestor_2(root, p, q *TreeNode) *TreeNode {
	path1 := make([]*TreeNode, 0)
	path2 := make([]*TreeNode, 0)
	DFS_2(root, p, &path1)
	DFS_2(root, q, &path2)
	var last *TreeNode
	for i := 0; i < len(path1) && i < len(path2); i++ {
		if path1[len(path1)-1-i].Val == path2[len(path2)-1-i].Val {
			last = path1[len(path1)-1-i]
			continue
		}
		break
	}
	return last
}

func test2() {
	path1 := make([]*TreeNode, 0)
	path2 := make([]*TreeNode, 0)
	DFS_2(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{0, nil, nil},
			Right: &TreeNode{8, nil, nil},
		},
	}, &TreeNode{2, nil, nil}, &path1)

	DFS_2(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{0, nil, nil},
			Right: &TreeNode{8, nil, nil},
		},
	}, &TreeNode{8, nil, nil}, &path2)

	fmt.Println("result:")
	for i := 0; i < len(path1); i++ {
		fmt.Print(path1[i].Val, "-")
	}
	fmt.Println()
	for i := 0; i < len(path2); i++ {
		fmt.Print(path2[i].Val, "-")
	}
	// 找相交部分
	var last *TreeNode
	for i := 0; i < len(path1) && i < len(path2); i++ {
		if path1[len(path1)-1-i].Val == path2[len(path2)-1-i].Val {
			last = path1[len(path1)-1-i]
			continue
		}
		break
	}
	if last != nil {
		fmt.Println("find:", last.Val)
	}

}

func DFS_2(root *TreeNode, p *TreeNode, path *[]*TreeNode) bool {
	// terminator
	if root == nil {
		return false
	}
	if root.Val == p.Val {
		*path = append(*path, root)
		return true
	}
	res1 := DFS_2(root.Left, p, path)
	if res1 {
		*path = append(*path, root)
		return true
	}
	res2 := DFS_2(root.Right, p, path)
	if res2 {
		*path = append(*path, root)
		return true
	}
	return false
}

// 使用Map存储
func lowestCommonAncestor_3(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return root
	}
	stack := []*TreeNode{root}
	parentMap := make(map[*TreeNode]*TreeNode) // 存储parent
	parentMap[root] = nil
	var path1Start *TreeNode
	var path2Start *TreeNode

	// 迭代需要两层for循环
	for len(stack) > 0 {
		popRes := stack[len(stack)-1] //  获取最后一个
		stack = stack[:len(stack)-1]
		if popRes.Val == p.Val {
			path1Start = popRes
		}
		if popRes.Val == q.Val {
			path2Start = popRes
		}
		if popRes.Left != nil {
			stack = append(stack, popRes.Left)
			parentMap[popRes.Left] = popRes
		}
		if popRes.Right != nil {
			stack = append(stack, popRes.Right)
			parentMap[popRes.Right] = popRes
		}
	}

	if path1Start == nil || path2Start == nil {
		return nil
	}

	// 找两个链表的重叠部分
	mapTemp := make(map[*TreeNode]bool) // 存储其中一条路径
	for path1Start != nil {
		mapTemp[path1Start] = true
		path1Start = parentMap[path1Start]
	}

	for path2Start != nil {
		if _, ok := mapTemp[path2Start]; ok {
			return path2Start
		}
		path2Start = parentMap[path2Start]
	}
	return nil
}
