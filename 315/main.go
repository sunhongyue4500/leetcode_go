package main

import "fmt"

// 1. 暴力
// 2. 普通分治，时间复杂度O(N^2)，空间复杂度O(N)
// 3. 归并排序+索引数组，时间复杂度O(NlgN)，空间复杂度O(N)
func main() {
	fmt.Println(countSmaller([]int{3, 2, 1}))
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
	fmt.Println(countSmaller([]int{5}))
	fmt.Println(countSmaller([]int{}))

	fmt.Println(countSmaller2([]int{3, 2, 1}))
	fmt.Println(countSmaller2([]int{5, 2, 6, 1}))
	fmt.Println(countSmaller2([]int{5}))
	fmt.Println(countSmaller2([]int{}))
}

func countSmaller(nums []int) []int {
	res := make([]int, len(nums))
	helper(nums, 0, len(nums)-1, &res)
	return res
}

func helper(nums []int, left, right int, temp *[]int) {
	// terminator
	if left >= right {
		return
	}
	// prepare data
	mid := left + (right-left)>>1
	// conquer sub problems
	helper(nums, left, mid, temp)
	helper(nums, mid+1, right, temp)
	// resolve result
	// left-right
	for i := left; i <= mid; i++ {
		for j := mid + 1; j <= right; j++ {
			if nums[j] < nums[i] {
				(*temp)[i]++
			}
		}
	}
}

// 归并排序+索引数组
func countSmaller2(nums []int) []int {
	tmp := make([]int, len(nums))   // 临时数组
	res := make([]int, len(nums))   // 结果数组
	index := make([]int, len(nums)) // 索引数组
	for i := 0; i < len(nums); i++ {
		index[i] = i // 初始化索引数组
	}
	helper2(nums, 0, len(nums)-1, index, tmp, res)
	return res
}

func helper2(nums []int, left, right int, index, temp, res []int) {
	// terminator
	if left >= right {
		return
	}
	// prepare data
	mid := left + (right-left)>>1
	// conquer sub problems
	helper2(nums, left, mid, index, temp, res)
	helper2(nums, mid+1, right, index, temp, res)
	// resolve result
	// left-right
	merge2(nums, left, mid, right, index, temp, res)
}

// index是索引数组，tmp是额外操作用的数组
func merge2(arr []int, low, mid, high int, index, tmp, res []int) {
	i, j, k := low, mid+1, low
	for i <= mid && j <= high {
		if arr[index[i]] <= arr[index[j]] {
			tmp[k] = index[j]
			k, j = k+1, j+1
		} else {
			res[index[i]] += high - j + 1 // 关键代码，计数
			tmp[k] = index[i]
			k, i = k+1, i+1
		}
	}
	for i <= mid {
		tmp[k] = index[i]
		k, i = k+1, i+1
	}
	for j <= high {
		tmp[k] = index[j]
		k, j = k+1, j+1
	}
	copy(index[low:high+1], tmp[low:high+1])
}
