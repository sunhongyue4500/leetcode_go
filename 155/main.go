package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(3)
	obj.Push(1)
	obj.Push(5)
	obj.Push(8)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()
	fmt.Println(param_3, param_4)
}

// 执行用时 :20 ms, 在所有 golang 提交中击败了90.60%的用户
// 内存消耗 :7.5 MB, 在所有 golang 提交中击败了47.56%的用户
type MinStack struct {
	stack1 []int // 数据栈
	stack2 []int //
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {
	this.stack1 = append(this.stack1, x)
	if len(this.stack2) == 0 || x <= this.stack2[len(this.stack2)-1] {
		this.stack2 = append(this.stack2, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack1) > 0 && len(this.stack2) > 0 {
		if this.stack1[len(this.stack1)-1] == this.stack2[len(this.stack2)-1] {
			this.stack2 = this.stack2[:len(this.stack2)-1]
		}
		this.stack1 = this.stack1[:len(this.stack1)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack1) > 0 {
		return this.stack1[len(this.stack1)-1]
	}
	return 1<<63 - 1
}

func (this *MinStack) GetMin() int {
	if len(this.stack2) > 0 {
		return this.stack2[len(this.stack2)-1]
	}
	return 1<<63 - 1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
