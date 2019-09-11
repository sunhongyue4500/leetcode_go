package main

import (
	"testing"
)

type Node struct {
	Val  int
	Next *Node
}

func test(head *Node) bool {
	// 使用快慢指针
	slow := head
	fast := head
	for {
		if slow == nil || fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
}

func Test141(t *testing.T) {
	// var endNode Node
	var node1 = Node{Val: 1}
	var node2 = Node{Val: 2}
	// var node3 = Node{Val: 3}
	// var node4 = Node{Val: 4}
	// var node5 = Node{Val: 5}
	// node1 := Node{1, &Node{2, &endNode}}
	node1.Next = &node2
	node2.Next = &node1
	// node3.Next = &node1
	// node4.Next = &node5
	// node5.Next = &node5
	t.Log(test(&node1))
}
