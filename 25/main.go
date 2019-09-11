package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1 := ListNode{1, nil}
	node2 := ListNode{2, &node1}
	node3 := ListNode{3, &node2}
	node4 := ListNode{4, &node3}
	node5 := ListNode{5, &node4}
	node6 := ListNode{6, &node5}
	node6.Next = nil
	//var node7 *ListNode

	res := reverseKGroup_2(&node5, 4)
	printListNode(res)

	//temp := reverseList(&node4)

	//temp2 := reverseList_2(&node4)
	//printListNode(temp2)

	//temp3 := reverseList_it(&node4)
	//printListNode(temp3)
}

// 递归实现
func reverseKGroup(head *ListNode, k int) *ListNode {
	start := head
	last := head
	// Terminator
	for i := 0; i < k; i++ {
		if head == nil {
			return start
		}
		last = head
		head = head.Next
	}
	// Drill Down
	nextHead := reverseKGroup(head, k)
	// 反转
	revEnd := reverse(start, k)
	revEnd.Next = nextHead
	return last
}

// 常数空间只能是迭代
func reverseKGroup_2(head *ListNode, k int) *ListNode {
	dummy := &ListNode{-1, head}
	pre := dummy
	cur := head
	for i := 1; cur != nil; i++ {
		cur = cur.Next
		if i%k == 0 {
			pre = reverseFromPre2Next(pre, cur)
		}
	}
	return dummy.Next
}

// 多写
// 反转pre和next之间的节点，并返回中间节点的最后一个节点
func reverseFromPre2Next(pre *ListNode, next *ListNode) *ListNode {
	// 因为last反转之后是最后一个，所以last不动
	last := pre.Next
	// cur 作为游标
	cur := last.Next
	for cur != next {
		last.Next = cur.Next // 往后指
		cur.Next = pre.Next  // cur往前指，反转
		pre.Next = cur       // 反转后，pre指向第一个
		cur = last.Next      // cur往后移
	}
	return last
}

// 反转pre和next之间的节点，并返回中间节点的最后一个节点
func reverseFromPre2Next_2(pre *ListNode, next *ListNode) *ListNode {
	// 因为last反转之后是最后一个，所以last不动

}

// 反转链表
func reverse(head *ListNode, k int) *ListNode {
	cur := head
	var pre *ListNode
	var next *ListNode
	for i := 0; i < k && cur != nil; i++ {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return head
}

func printListNode(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
