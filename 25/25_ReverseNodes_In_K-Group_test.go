package main

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代思路
func reverseKGroup(head *ListNode, k int) *ListNode {
	flag := head
	counter := 1
	var pre *ListNode
	var cur = head
	var next = head.Next

	for flag.Next != nil {
		flag = flag.Next
		counter++
		if counter == k {
			flag = flag.Next
			// 反转之前的节点
			cur.Next = flag
			for cur != flag {
				pre = cur
				cur = next
				next = next.Next
				cur.Next = pre
			}
			counter = 1
		}
	}

	// // 主逻辑
	// cur.Next = pre
	// pre = cur
	// cur = next
	// next = next.Next

	return nil
}

// 以pre作为头，next作为尾，反转中间的节点，返回反转中间节点的最后一个节点
// pre和next中间至少要有一个节点
func reverseKGroupOnce(pre *ListNode, next *ListNode) *ListNode {
	// 因为last反转之后是最后一个，所以last不动
	last := pre.Next
	if last == next { // 中间没有节点
		return pre
	}
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

func reverseKGroup2(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	// k == 1
	// 添加一个dummy node
	dummy := new(ListNode)
	dummy.Val = -1
	dummy.Next = head
	pre := dummy
	cur := head
	for i := 1; cur != nil; i++ {
		cur = cur.Next
		if i%k == 0 {
			pre = reverseKGroupOnce(pre, cur) // 满足k个才反转
		}
	}
	fmt.Println("AAA", pre.Val)
	fmt.Println("CUR", cur == nil)
	// 不足的再反转
	// reverseKGroupOnce(pre, cur)	// 这句控制不足的是否反转
	return dummy.Next
}

func TestReverserKGroup(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 1
	node2 := new(ListNode)
	node2.Val = 2
	node3 := new(ListNode)
	node3.Val = 3
	node4 := new(ListNode)
	node4.Val = 4

	node5 := new(ListNode)
	node5.Val = 5

	node1.Next = node2
	// node2.Next = nil
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = nil

	node := reverseKGroup2(node1, 3)
	for node != nil {
		t.Log("result:", node.Val)
		node = node.Next
	}
}

func TestReverserKGroupOnce(t *testing.T) {
	node1 := new(ListNode)
	node1.Val = 1
	node2 := new(ListNode)
	node2.Val = 2
	node3 := new(ListNode)
	node3.Val = 3
	node4 := new(ListNode)
	node4.Val = 4

	// node1.Next = node2
	node1.Next = nil
	// node2.Next = nil
	node2.Next = node3
	node3.Next = node4
	node4.Next = nil

	dummyNode := new(ListNode)
	dummyNode.Val = -1
	dummyNode.Next = node1

	fmt.Println("start")
	result := reverseKGroupOnce(dummyNode, nil)
	fmt.Println("end")
	t.Log("out:", result.Val)

	temp := dummyNode
	for temp != nil {
		t.Log("rrr:", temp.Val)
		temp = temp.Next
	}
}
