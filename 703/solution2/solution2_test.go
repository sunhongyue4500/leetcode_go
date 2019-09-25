package solution2

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
	"testing"
)

// 不用内置库的解法

type KthLargest struct {
	heap MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	// assert k >= 1
	// assert len(nums) >= k-1
	// nums = [] while k = 1 is possible

	n := len(nums)
	if k < n {
		n = k
	}

	MIN_INT := -int(^uint(0) >> 1) -1
	fmt.Println(MIN_INT)
	fmt.Println(intsets.MinInt)
	array := make([]int, k)
	for i := 0; i< k; i++ {
		array[i] = MIN_INT
	}

	heap := MinHeap{
		Array: array,
		N: n,
	}

	for _, num := range nums {
		heap.Add(num)
	}

	return KthLargest{heap: heap}

}


func (this *KthLargest) Add(val int) int {
	this.heap.Add(val)
	return this.heap.Array[0]

}

type MinHeap struct {
	Array []int
	N int
}

func (heap *MinHeap) Add(v int) {

	if heap.N == len(heap.Array) {
		if v < heap.Array[0] {
			return
		}
	} else {
		heap.N += 1
	}

	heap.Array[0] = v

	heap.balance(0)
}

func (heap *MinHeap) balance(index int) {
	left := index * 2 +1
	right := index * 2 + 2
	n := len(heap.Array)
	if left >= n {
		// reach the leaf
		return
	}

	minIndex := index
	if heap.Array[minIndex] > heap.Array[left] {
		minIndex = left
	}
	if right < n {
		if heap.Array[minIndex] > heap.Array[right] {
			minIndex = right
		}
	}
	if minIndex == index {
		return
	}

	heap.Array[index], heap.Array[minIndex] = heap.Array[minIndex], heap.Array[index]

	// 递归进行平衡
	heap.balance(minIndex)

}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func TestSolution2(t *testing.T) {
	obj := Constructor(2, []int{3, 8, 2})
	param_1 := obj.Add(1)
	param_1 = obj.Add(9)
	fmt.Println(param_1)
}