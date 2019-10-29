package main

import (
	"fmt"
	"sort"
)

func main() {
	test := []int{2, 2, 1, 1, 1, 2, 2}
	test2 := []int{1, 2, 3}
	fmt.Println(majorityElement(test))
	fmt.Println(majorityElement_2(test))
	fmt.Println()
	fmt.Println(majorityElement_3(test))
	// 如果不存在众数，第三个解法有问题
	fmt.Println(majorityElement_3(test2))
}

//  直接排序
func majorityElement(nums []int) int {
	temp := nums
	sort.Ints(temp)
	fmt.Println(temp)
	return temp[len(temp)/2]
}

// 用map来记录出现的次数
func majorityElement_2(nums []int) int {
	mapTemp := make(map[int]int)
	length := len(nums) / 2
	for _, val := range nums {
		mapTemp[val] += 1
		if mapTemp[val] > length {
			return val
		}
	}
	return nums[0]
}

// 分治 divide and conquer
// 对左边和右边分别求众数，
func majorityElement_3(nums []int) int {
	return helper1(nums, 0, len(nums)-1)
}

func countInRange(nums []int, num int, li int, hi int) int {
	count := 0
	for i := li; i <= hi; i++ {
		if nums[i] == num {
			count++
		}
	}
	return count
}

// 计算从li到hi的众数
func helper1(nums []int, li int, hi int) int {
	// terminator
	if li == hi {
		return nums[li]
	}

	mid := (hi-li)/2 + li

	// drill down
	left := helper1(nums, li, mid)
	right := helper1(nums, mid+1, hi)

	if left == right {
		return left
	}

	// 计算出现的次数
	leftcount := countInRange(nums, left, li, hi)
	rightcount := countInRange(nums, right, li, hi)

	if leftcount > rightcount {
		return left
	}
	return right
}
