package main

// quicksort 排序实现
func QSort(array *[]int, start int, end int) []int {
	if start >= end || start < 0 || end < 0 {
		return *array
	}
	// 挖坑添数
	x := start
	y := end
	target := start
	// 从后往前找，找第一个比target小的
	targetVal := (*array)[target]

	for x < y {
		// 从后往前找第一个比它小
		for y > x && (*array)[y] >= targetVal {
			y--
		}
		// 交换
		(*array)[y], (*array)[target] = (*array)[target], (*array)[y]
		for y > x && (*array)[x] <= targetVal {
			x++
		}
		(*array)[x], (*array)[y] = (*array)[y], (*array)[x]
		target = x
	}
	QSort(array, start, x-1)
	QSort(array, x+1, end)
	//res := []int{}
	return *array
}

func QSort2(array []int, left int, right int) []int {
	if left >= right {
		return array
	}
	start := left
	end := right
	partion := array[start]

	for start != end {
		// 从后往前找
		for end > start && array[end] >= partion {
			end--
		}
		array[start] = array[end]
		for end > start && array[start] <= partion {
			start++
		}
		array[end] = array[start]
		array[start] = partion
	}

	QSort2(array, left, start-1)
	QSort2(array, start+1, right)
	return array
}