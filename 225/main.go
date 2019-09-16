package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	param_2 := obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Empty()
	fmt.Println(param_2, param_3, param_4)
	obj.Push(4)
	obj.Push(5)
	param_2 = obj.Pop()
	param_2 = obj.Pop()
	param_2 = obj.Pop()
	param_3 = obj.Top()
	param_4 = obj.Empty()
	fmt.Println(param_2, param_3, param_4)
}


type Queue struct {
	data []int
}

func NewQueue(len int) *Queue{
	return &Queue{make([]int, len)}
}

func (queue *Queue)push(ele int) {
	queue.data = append(queue.data, ele)
}

func (queue *Queue)pop() int {
	if queue.empty() {
		return -1
	}
	res := queue.data[0]
	queue.data = queue.data[1:]
	return res
}

func (queue *Queue)empty() bool {
	return queue.size() == 0
}

func (queue *Queue)size() int {
	return len(queue.data)
}

func (queue *Queue)peek() int {
	if queue.empty() {
		return -1
	}
	res := queue.data[0]
	return res
}

// 使用两个队列
//type MyStack struct {
//	q0 *Queue
//	q1 *Queue
//}
//
//
///** Initialize your data structure here. */
//func Constructor() MyStack {
//	return MyStack{
//		q0:NewQueue(0),
//		q1:NewQueue(0),
//	}
//}
//
//
///** Push element x onto stack. */
//func (this *MyStack) Push(x int)  {
//	this.q0.push(x)
//}
//
//
///** Removes the element on top of the stack and returns that element. */
//func (this *MyStack) Pop() int {
//	if this.q0.empty() {
//		return -1
//	}
//	for this.q0.size() > 1 {
//		this.q1.push(this.q0.pop())
//	}
//	res := this.q0.pop()
//	// swap
//	temp := this.q0
//	this.q0 = this.q1
//	this.q1 = temp
//	return res
//}
//
//
///** Get the top element. */
//func (this *MyStack) Top() int {
//	if this.q0.empty() {
//		return -1
//	}
//	for this.q0.size() > 1 {
//		this.q1.push(this.q0.pop())
//	}
//	res := this.q0.peek()
//	this.q1.push(this.q0.pop())
//	// swap
//	temp := this.q0
//	this.q0 = this.q1
//	this.q1 = temp
//	return res
//}
//
//
///** Returns whether the stack is empty. */
//func (this *MyStack) Empty() bool {
//	return this.q0.empty()
//}


// 用一个队列来实现
type MyStack struct {
	q *Queue
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{NewQueue(0)}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	// push的时候把队列颠倒
	size := this.q.size()
	this.q.push(x)
	for i:=0; i<size; i++ {
		this.q.push(this.q.pop())
	}
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.q.pop()
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q.peek()
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q.empty()
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */