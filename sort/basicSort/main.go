package main

import "fmt"

// 1. 选择排序，两遍for选最小的，交换。不稳定
// 2. 插入排序，两遍for，移动数据，插入到指定位置。稳定
// 3. 冒泡排序，两遍for，向上冒最大的。稳定
// 4. 归并排序，递归分治，原地合并两个有序数组。稳定
// 5. 快速排序，选出一个轴心，调配出两个数组，分而治之。不稳定
// 6. 堆排序，通过构造二叉堆获取堆顶元素实现排序。不稳定
func main() {
	test := []int{6, 4, 3, -4, 2, 8, 1}
	backNums := make([]int, len(test))
	copy(backNums, test)
	backHeapNums := make([]int, len(test))
	copy(backHeapNums, test)
	fmt.Println("选择排序")
	fmt.Println(selectionSort(test))
	fmt.Println(test)
	fmt.Println("插入排序")
	fmt.Println(insertionSort(test))
	fmt.Println(test)
	fmt.Println("冒泡排序")
	fmt.Println(bubbleSort(test))
	fmt.Println(test)
	temp := []int{1, 3, 5, 2, 4, 6}
	//merge(temp, 0, 2, 5)
	//fmt.Println(temp)
	//mergeHelper(temp, 0, len(temp)-1)
	fmt.Println("归并排序")
	fmt.Println(mergeSort(test))
	fmt.Println(mergeSort(temp))
	fmt.Println(test)
	fmt.Println("快速排序")
	quickSort(backNums, 0, len(backNums)-1)
	fmt.Println(backNums)
	fmt.Println(test)
	fmt.Println("堆排序")
	fmt.Println(heapSort(backHeapNums))
	fmt.Println(test)
}

// 通过看动画可写出代码

/**
选择排序，不稳定？
序列5 8 5 2 9，我们知道第一遍选择第1个元素5会和2交换，那么原序列中2个5的相对前后顺序就被破坏了，所以选择排序不是一个稳定的排序算法。
两遍for，定住一个，另一遍for找最小的
*/
func selectionSort(nums []int) []int {
	backNums := make([]int, len(nums))
	copy(backNums, nums)
	min := 0
	for i := 0; i < len(backNums); i++ {
		min = i
		for j := i + 1; j < len(backNums); j++ {
			if backNums[j] < backNums[min] {
				min = j
			}
		}
		backNums[i], backNums[min] = backNums[min], backNums[i]
	}
	return backNums
}

/**
插入排序，两遍for
*/
func insertionSort(nums []int) []int {
	backNums := make([]int, len(nums))
	copy(backNums, nums)
	for i := 0; i < len(backNums); i++ {
		min := i
		temp := backNums[i]
		for j := i - 1; j >= 0; j-- {
			backNums[j+1] = backNums[j] // 后移
			// 找到第一个比这个小
			if backNums[j] < temp {
				break
			}
			min = j
		}
		backNums[min] = temp
	}
	return backNums
}

/**
冒泡排序，两遍for
*/
func bubbleSort(nums []int) []int {
	backNums := make([]int, len(nums))
	copy(backNums, nums)
	for i := len(backNums) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if backNums[j] > backNums[j+1] {
				backNums[j], backNums[j+1] = backNums[j+1], backNums[j]
			}
		}
	}
	return backNums
}

func mergeSort(nums []int) []int {
	backNums := make([]int, len(nums))
	copy(backNums, nums)
	mergeSortHelper(backNums, 0, len(backNums)-1)
	return backNums
}

// 递归函数
func mergeSortHelper(nums []int, left, right int) {
	// terminator
	if left >= right {
		return
	}
	// drill down
	mid := (left + right) >> 1
	mergeSortHelper(nums, left, mid)
	mergeSortHelper(nums, mid+1, right)
	merge(nums, left, mid, right)
}

/**
从left到mid是有序的，从mid+1到right是有序的
*/
func merge(nums []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	// 比较起始部分
	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp[k], k, i = nums[i], k+1, i+1
		} else {
			temp[k], k, j = nums[j], k+1, j+1
		}
	}
	// 处理剩余的
	for i <= mid {
		temp[k], k, i = nums[i], k+1, i+1
	}
	for j <= right {
		temp[k], k, j = nums[j], k+1, j+1
	}
	// 将临时数组，拷贝回原数组
	for p := 0; p < len(temp); p++ {
		nums[left+p] = temp[p]
	}
}

/**
快排
*/
func quickSort(nums []int, left, right int) {
	if right <= left {
		return
	}
	pivot := partion(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

// 比pivot都放左，大的都放右
func partion(nums []int, left, right int) int {
	counter, pivot := left, right
	for i := left; i < right; i++ {
		if nums[i] < nums[pivot] { // 小的都从左边开始放
			nums[i], nums[counter] = nums[counter], nums[i]
			counter++
		}
	}
	nums[counter], nums[pivot] = nums[pivot], nums[counter]
	return counter
}

func heapSort(nums []int) []int {
	res := []int{}
	if len(nums) == 0 {
		return nil
	}
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapfiy(nums, len(nums), i)
	}

	for i := len(nums) - 1; i >= 0; i-- {
		res = append(res, nums[0])
		nums[0], nums[i] = nums[i], nums[0]
		heapfiy(nums, i, 0)
	}
	return res
}

/**
堆化，构造小顶堆(或大顶堆)
i:对位置i做堆化
length:到哪截止
*/
func heapfiy(nums []int, length, i int) {
	left, right := 2*i+1, 2*i+2
	largest := i
	if left < length && nums[left] < nums[largest] {
		largest = left
	}
	if right < length && nums[right] < nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapfiy(nums, length, largest)
	}
}
