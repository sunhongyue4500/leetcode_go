package deque

// 用slice实现
type Deque struct {
	data []int
}

func (deque *Deque)AddFirst(ele int)  {
	deque.data = append([]int{ele}, deque.data...)
}

func (deque *Deque)AddLast(ele int) {
	deque.data = append(deque.data, ele)
}

func (deque *Deque)RemoveFirst() int {
	res := deque.data[0]
	deque.data = deque.data[1:]
	return res
}

func (deque *Deque)RemoveLast() int {
	res := deque.data[len(deque.data) - 1]
	deque.data = deque.data[:len(deque.data) - 1]
	return res
}