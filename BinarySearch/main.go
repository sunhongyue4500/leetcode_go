package main

import "fmt"

func main() {
	arr := []int{1, 5, 18, 34, 45, 47, 55, 67, 69, 70, 71, 88, 89, 92, 100}
	target := 1
	fmt.Println(bs_iteration(arr, target))

	fmt.Println(bs_recursion(arr, target, 0, len(arr)-1))
}

// 迭代
func bs_iteration(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// 递归
func bs_recursion(arr []int, target int, left int, right int) int {
	// terminator
	if left > right {
		return -1
	}
	// process
	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	}
	// drill down
	if arr[mid] > target {
		return bs_recursion(arr, target, left, mid-1)
	} else {
		return bs_recursion(arr, target, mid+1, right)
	}
	// clear status
}
