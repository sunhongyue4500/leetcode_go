package main

import "fmt"

func main() {
	fmt.Println(mergeSortWithIndexArray([]int{3, 1, 2, 0}))
	fmt.Println(mergeSortWithIndexArray([]int{3, 5, 4, 1, 2, 0}))
}

// 返回索引数组，如nums：[3, 1, 2, 0], 结果就是[3, 1, 2, 0]
func mergeSortWithIndexArray(nums []int) []int {
	idx := make([]int, len(nums))
	for i := 0; i < len(idx); i++ {
		idx[i] = i
	}
	temp := make([]int, len(nums))
	mergeSort(nums, 0, len(nums)-1, temp, idx)
	return idx
}

// 归并排序+索引数组
func mergeSort(nums []int, left, right int, temp, idx []int) {
	if left >= right {
		return
	}
	mid := left + (right-left)>>1
	mergeSort(nums, left, mid, temp, idx)
	mergeSort(nums, mid+1, right, temp, idx)
	merge(nums, left, mid, right, temp, idx)
}

func merge(nums []int, left, mid, right int, temp []int, idx []int) {
	i, j, k := left, mid+1, left
	for i <= mid && j <= right {
		if nums[idx[i]] < nums[idx[j]] {
			temp[k], k, i = idx[i], k+1, i+1
		} else {
			temp[k], k, j = idx[j], k+1, j+1
		}
	}
	// 处理剩余的
	for i <= mid {
		temp[k], k, i = idx[i], k+1, i+1
	}
	for j <= right {
		temp[k], k, j = idx[j], k+1, j+1
	}
	// copy
	copy(idx[left:right+1], temp[left:right+1])
}
