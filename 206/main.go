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

	//temp := reverseList(&node4)
	//printListNode(temp)

	temp2 := reverseList_2(&node4)
	printListNode(temp2)

	//temp3 := reverseList_it(&node4)
	//printListNode(temp3)
}

func printListNode(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

// 迭代反转链表
//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//	var pre *ListNode
//	cur := head
//	next := cur.Next
//	for next != nil {
//		// 反转
//		cur.Next = pre
//		pre = cur
//		cur = next
//		next = next.Next
//	}
//	cur.Next = pre
//	return cur
//}

// 迭代优化
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var cur *ListNode
	for head != nil {
		cur = head
		head = head.Next
		cur.Next = pre
		pre = cur
	}
	return pre
}

func reverseList_it(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	var next *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 递归解法
/*
var reverseList = function (head) {
    if (!head || !(head.next)) {
        return head
    }
    let nextNode = head.next
    // 找到最后一个用于返回
    let result = reverseList(head.next)
    nextNode.next = head
    head.next = undefined
    return result
}
*/
func reverseList_2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList_2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
