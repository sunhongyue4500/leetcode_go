package re_solution2

import (
	"fmt"
	"testing"
)

type KthLargest struct {
	heap MinHeap
}

type MinHeap struct {
	array []int
	N int
}

func Constructor(k int, nums []int) KthLargest {
	min := -int(^uint(0) >> 1) - 1
	//fmt.Println(min)
	arr := make([]int, k)
	for i:=0; i<k; i++ {
		// 初始化为最小值
		arr[i] = min
	}

	minHeap := MinHeap{arr, k}
	heap := KthLargest{minHeap}

	for _, val := range nums {
		heap.Add(val)
	}

	return heap
}

func (minHeap *MinHeap)add(val int) {
	if minHeap.N == len(minHeap.array) {
		if val < minHeap.array[0] {
			return
		}
	} else {
		minHeap.N += 1
	}
	// 把最小的踢出去，换上新的值
	minHeap.array[0] = val
	// 平衡堆
	minHeap.balence(0)
}

// 递归来实现下沉
func (minHeap *MinHeap)balence(idx int) {
	// 通过下沉来平衡
	left := 2 * idx + 1
	right := left + 1
	// 无叶子
	if left > minHeap.N - 1 {
		return
	}
	swapIdx := -1
	// 有左
	if left < minHeap.N && right > minHeap.N - 1 {
		if minHeap.array[left] < minHeap.array[idx] {
			swapIdx = left
		} else {
			return
		}
	}
	// 两个都有
	if right < minHeap.N {
		minTemp := -1
		if minHeap.array[left] < minHeap.array[right] {
			minTemp = left
		} else {
			minTemp = right
		}
		if minHeap.array[minTemp] < minHeap.array[idx] {
			swapIdx = minTemp
		} else {
			return
		}
	}
	// 交换
	minHeap.array[idx], minHeap.array[swapIdx] = minHeap.array[swapIdx], minHeap.array[idx]
	minHeap.balence(swapIdx)
}

func (this *KthLargest) Add(val int) int {
	this.heap.add(val)
	return this.heap.array[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func TestReSolution2(t *testing.T) {
	obj := Constructor(2, []int{3, 8, 2})
	param_1 := obj.Add(1)
	fmt.Println(param_1)
	param_1 = obj.Add(9)
	fmt.Println(param_1)
}