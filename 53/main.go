package main

import "fmt"

func main() {
	test := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(test))
}

// 暴力法：所有可能的方式都尝试一遍
// 使用DP
// 使用分治法

// 分治
func maxSubArray(nums []int) int {
	return help(nums, 0, len(nums)-1)
}

func help(nums []int, l, r int) int {
	if l > r {
		return -(1 << 63)
	}
	m := l + (r-l)/2
	ml, mr := 0, 0
	// 分治
	lMax := help(nums, l, m-1)
	rMax := help(nums, m+1, r)
	for i, sum := m-1, 0; i >= l; i-- {
		sum += nums[i]
		ml = max(sum, ml)
	}
	for i, sum := m+1, 0; i <= r; i++ {
		sum += nums[i]
		mr = max(sum, mr)
	}
	return max(max(lMax, rMax), ml+mr+nums[m])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 暴力，时间复杂度O(N^2)，空间复杂度O(1)
func maxSubArray2(nums []int) int {
	res := nums[0]
	temp := 0
	for i := 0; i < len(nums); i++ {
		temp = 0
		for j := i; j < len(nums); j++ {
			temp += nums[j]
			res = max(temp, res)
		}
	}
	return res
}
