package main

import "fmt"

func main() {
	arr := []int{1, 5, 18, 34, 45, 47, 2, 69, 9, 70, 71, 55, 67, 88, 89, 92, 100}
	qsort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func qsort(arr []int, left int, right int) {
	start := left
	end := right
	partion := arr[start]
	// terminator
	if start >= end {
		return
	}
	// process
	for start != end {
		// 从右往左找，找第一个比partion小的
		for end > start && arr[end] >= partion {
			end--
		}
		arr[start] = arr[end]
		// 从左向右找，找第一个比partion大的
		for end > start && arr[start] <= partion {
			start++
		}
		arr[end] = arr[start]
		arr[start] = partion
	}
	// drill down
	qsort(arr, left, start-1)
	qsort(arr, start+1, right)
	// clear status
}
