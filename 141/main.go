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
	res := hasCycle_2(&node5)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用快慢指针
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head.Next, head.Next.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// 使用map来存储
func hasCycle_2(head *ListNode) bool {
	temp := head
	mmap := make(map[*ListNode]bool)
	for temp != nil {
		if _, ok := mmap[temp]; ok {
			return true
		}
		mmap[temp] = true
		temp = temp.Next
	}
	return false
}
