package stack_queue


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