package main

import "fmt"

func main() {
	test1 := []int{1, 5, -1, 0, 10}
	test2 := []int{0, -1, -5, 0, -4}
	fmt.Println(solution1(test1))
	fmt.Println(solution1(test2))
	fmt.Println(solution2(test1))
	fmt.Println(solution2(test2))
}

// dp，其实就是Kadane’s Algorithm算法
// dp[i]表示截止到第i个元素的最大连续
func solution1(array []int) int {
	dp := make([]int, len(array))
	dp[0] = array[0]
	res := dp[0]
	for i := 1; i < len(array); i++ {
		dp[i] = max(array[i], dp[i-1]+array[i])
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// DP，状态压缩
func solution2(array []int) int {
	if len(array) == 0 {
		return 0
	}
	cur := array[0]
	res := cur
	for i := 1; i < len(array); i++ {
		cur = max(array[i], cur+array[i])
		if res < cur {
			res = cur
		}
	}
	return res
}

// 使用分治法
func solution3(nums []int) int {
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
