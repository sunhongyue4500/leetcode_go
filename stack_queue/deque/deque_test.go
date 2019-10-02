package deque

import "testing"

func TestDeque(t *testing.T) {
	deque := Deque{[]int{2, 3, 4, 5}}
	deque.AddFirst(6)
	deque.AddLast(10)
	t.Log(deque.data)
	t.Log(deque.RemoveFirst())
	t.Log(deque.RemoveLast())
	t.Log(deque.data)
}