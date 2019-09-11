package _32

import (
	"errors"
	"testing"
)

type Stack struct {
	StackSlice []int
}

func (stack *Stack) push(ele int) {
	stack.StackSlice = append(stack.StackSlice, ele)
}

func (stack *Stack) pop() (int, error) {
	if stack.empty() {
		return -1, errors.New("empty")
	}
	temp := stack.StackSlice[len(stack.StackSlice)-1]
	// reslice
	stack.StackSlice = stack.StackSlice[0 : len(stack.StackSlice)-1]
	return temp, nil
}

func (stack *Stack) size() int {
	return len(stack.StackSlice)
}

func (stack *Stack) peek() (int, error) {
	if stack.empty() {
		return -1, errors.New("empty")
	}
	return stack.StackSlice[len(stack.StackSlice)-1], nil
}

func (stack *Stack) empty() bool {
	if stack.size() == 0 {
		return true
	}
	return false
}

type MyQueue struct {
	stack       *Stack
	back_statck *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	myQ := *(new(MyQueue))
	myQ.stack = &Stack{make([]int, 0)}
	myQ.back_statck = &Stack{make([]int, 0)}
	return myQ
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack.push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	for this.stack.size() > 1 {
		if val, err := this.stack.pop(); err == nil {
			this.back_statck.push(val)
		} else {
			return -99
		}
	}
	result, _ := this.stack.pop()
	for !this.back_statck.empty() {
		temp, _ := this.back_statck.pop()
		this.stack.push(temp)
	}
	return result
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	for this.stack.size() > 1 {
		if val, err := this.stack.pop(); err == nil {
			this.back_statck.push(val)
		} else {
			return -1
		}
	}
	result, _ := this.stack.peek()

	for !this.back_statck.empty() {
		temp, _ := this.back_statck.pop()
		this.stack.push(temp)
	}
	return result
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack.empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func TestQueueWithStack(t *testing.T) {
	obj := Constructor()
	t.Log("res", obj.Pop())
	t.Log("res", obj.Pop())
	t.Log("res", obj.Pop())
	obj.Push(10)
	t.Log("res", obj.Pop())
	obj.Push(4)
	obj.Push(3)
	t.Log("res", obj.stack)
	t.Log("res", obj.Pop())
	obj.Push(2)

	t.Log("res1", obj.Peek())
	t.Log(obj.Empty())
	t.Log("res", obj.stack)
	t.Log(obj.Pop())
	t.Log("res1", obj.Peek())
	t.Log("res1", obj.Peek())
	t.Log("res1", obj.Peek())
	t.Log(obj.Pop())

	t.Log(obj.Pop())
	t.Log(obj.Pop())
}
