package main

import "fmt"

func main() {
	test := []int{1, 2, 3}
	fmt.Println(subsets(test))
}

// 1. 暴力，枚举所有可能
// 2. 分治
func subsets(nums []int) [][]int {
	res := [][]int{}
	list := []int{}
	helper(nums, &list, &res, 0)
	return res
}

// 每一层进去都放或者不放
func helper(nums []int, list *[]int, res *[][]int, level int) {
	// terminator
	if level >= len(nums) {
		temp := make([]int, len(*list))
		copy(temp, *list)
		*res = append(*res, temp)
		return
	}
	// split problems
	// 不放
	helper(nums, list, res, level+1)
	// 放
	*list = append(*list, nums[level])
	helper(nums, list, res, level+1)
	// clear status
	*list = (*list)[:len(*list)-1]
}
