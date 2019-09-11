package main

type Stack struct {
	data []byte
}

func (stack *Stack) Push(ele byte) {
	stack.data = append(stack.data, ele)
}

func (stack *Stack) Pop() (ele byte) {
	if len(stack.data) == 0 {
		return '0'
	}
	elem := stack.data[len(stack.data) - 1]
	stack.data = stack.data[0:len(stack.data) - 1]
	return elem
}

func (stack *Stack)Equal(stack2 *Stack) bool {
	// 先判断长度
	if len(stack.data) != len(stack2.data) {
		return false
	}
	// 再比较
	for idx, val := range stack.data {
		if val != stack2.data[idx] {
			return false
		}
	}
	return true
}
