package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	if h.Len() == 0 {
		return nil
	}
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() interface{} {
	if h.Len() == 0 {
		return nil
	}
	old := *h
	x := old[0]
	return x
}


type KthLargest struct {
	MyHeap *IntHeap
	MyK int
}


func Constructor(k int, nums []int) KthLargest {
	tempHeap := &IntHeap{}
	heap.Init(tempHeap)
	for i:=0; i< len(nums); i++ {
		if tempHeap.Len() < k {
			// 不满
			heap.Push(tempHeap, nums[i])
		} else {
			// 比较堆顶和接下来的元素
			ele := tempHeap.Peek().(int)
			if nums[i] > ele {
				heap.Pop(tempHeap)
				heap.Push(tempHeap, nums[i])
			}
		}
	}

	return KthLargest{
		MyHeap:tempHeap,
		MyK:k,
	}
}


func (this *KthLargest) Add(val int) int {
	if this.MyHeap.Len() < this.MyK {
		heap.Push(this.MyHeap, val)
	} else {
		ele := this.MyHeap.Peek().(int)
		if val > ele {
			heap.Pop(this.MyHeap)
			heap.Push(this.MyHeap, val)
		}
	}
	return this.MyHeap.Peek().(int)
}



// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

	fmt.Println()
	tempKthLargest := Constructor(3, []int{1000, 4,5,8,2,9, 10, 100})
	//for tempKthLargest.MyHeap.Len() > 0 {
	//	fmt.Printf("%d ", heap.Pop(tempKthLargest.MyHeap))
	//}
	fmt.Println()
	fmt.Println(tempKthLargest.Add(10000))

	Constructor(1, []int{})
	//fmt.Println(heap.Pop(temp2.MyHeap))
}
