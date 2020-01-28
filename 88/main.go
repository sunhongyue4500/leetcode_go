package main

import "fmt"

func main() {
	nums1 := []int{1, 3, 5, 0, 0, 0, 0}
	nums2 := []int{2, 4, 6, 8}
	//merge(nums1, 3, nums2, 4)
	merge2(nums1, 3, nums2, 4)
	fmt.Println(nums1)
}

/**
时间复杂度O(m+n)，空间复杂度O(m+n)
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m+n)
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] > nums2[j] {
			temp[k], k, j = nums2[j], k+1, j+1
		} else {
			temp[k], k, i = nums1[i], k+1, i+1
		}
	}
	for i < m {
		temp[k], k, i = nums1[i], k+1, i+1
	}
	for j < n {
		temp[k], k, j = nums2[j], k+1, j+1
	}
	for p := 0; p < len(temp); p++ {
		nums1[p] = temp[p]
	}
}

/**
从尾部开始合并，只需走一遍，时间复杂度O(m+n)，空间复杂度O(1)
*/
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	for k := m + n - 1; k >= 0; k-- {
		if i < 0 {
			// 处理剩余的
			nums1[k], j = nums2[j], j-1
			continue
		}
		if j < 0 {
			nums1[k], i = nums1[i], i-1
			continue
		}
		if nums1[i] > nums2[j] {
			nums1[k], i = nums1[i], i-1
		} else {
			nums1[k], j = nums2[j], j-1
		}
	}
}
