package main

import "fmt"

func main() {
	//var node *ListNode
	node1 := ListNode{1, nil}
	node2 := ListNode{2, &node1}
	node3 := ListNode{3, &node2}
	//node4 := ListNode{4, &node3}
	//node5 := ListNode{5, &node4}
	res := swapPairs(&node3)
	printListNode(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归解法
func swapPairs(head *ListNode) *ListNode {
	// Terminator
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	// Drill down
	head.Next = swapPairs(next.Next)
	// Process
	next.Next = head
	return next
}

func printListNode(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

// 迭代解法
func swapPairs_2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head.Next
	pre := &ListNode{-1, nil}
	cur := head
	var next *ListNode
	for cur != nil && cur.Next != nil {
		next = cur.Next
		// 反转
		cur.Next = next.Next
		next.Next = cur
		pre.Next = next
		// 向后移动
		pre = cur
		cur = cur.Next
	}
	return res
}
