package stack_queue

type Queue struct {
	data []int
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
	return len(queue.data) == 0
}

func (queue *Queue)peek() int {
	if queue.empty() {
		return -1
	}
	res := queue.data[0]
	return res
}