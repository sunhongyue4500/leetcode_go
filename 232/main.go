package _32

import "fmt"

type Stack struct {
	data []int
}

func NewStack(len int) *Stack {
	return &Stack{make([]int, len)}
}

func (stack *Stack)push(ele int) {
	stack.data = append(stack.data, ele)
}

func (stack *Stack)empty() bool {
	return len(stack.data) == 0
}

func (stack *Stack)pop() int {
	if stack.empty() {
		return -1
	}
	res := stack.data[len(stack.data) - 1]
	stack.data = stack.data[0:len(stack.data) - 1]
	return res
}

func (stack *Stack)peek() int {
	if stack.empty() {
		return -1
	}
	res := stack.data[len(stack.data) - 1]
	return res
}

type MyQueue struct {
	inputStack *Stack
	outputStack *Stack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		inputStack:NewStack(0),
		outputStack:NewStack(0),
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.inputStack.push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if !this.outputStack.empty() {
		return this.outputStack.pop()
	}
	for !this.inputStack.empty() {
		this.outputStack.push(this.inputStack.pop())
	}
	return this.outputStack.pop()
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if !this.outputStack.empty() {
		return this.outputStack.peek()
	}
	for !this.inputStack.empty() {
		this.outputStack.push(this.inputStack.pop())
	}
	return this.outputStack.peek()
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.inputStack.empty() && this.outputStack.empty()
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	obj.Push(1)
	//obj.Push(2)
	//obj.Push(3)
	param_2 := obj.Pop()
	param_2 = obj.Pop()
	param_3 := obj.Peek()
	param_4 := obj.Empty()
	obj.Push(4)
	obj.Push(5)
	fmt.Println(param_2, param_3, param_4)
	param_2 = obj.Pop()
	//param_3 = obj.Peek()
	//param_4 = obj.Empty()
	fmt.Println(param_2, param_3, param_4)
}