package main

import "fmt"

func main() {
	//var node *ListNode
	node1 := ListNode{1, nil}
	node2 := ListNode{2, &node1}
	node3 := ListNode{3, &node2}
	node4 := ListNode{4, &node3}

	node5 := ListNode{5, &node4}
	node1.Next = &node2
	res := detectCycle(&node5)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	// 先用快慢指针看是否有环
	// 复用这两个指针
	slow := head
	fast := head
	for {
		if slow == nil || slow.Next == nil || fast == nil || fast.Next == nil {
			// 无环
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// 有环
			break
		}
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
