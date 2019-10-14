package main

import (
	"fmt"
	"strconv"
)

func main() {
	array := []int{3, 8, 4, 9, 1, 5, 2}
	array2 := []int{3, 8, 4, 9, 1, 5, 2}
	fmt.Println(QSort(&array, 0, len(array) - 1))
	fmt.Println(QSort2(array2, 0, len(array2) - 1))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}))
	fmt.Println()
	fmt.Println(threeSum_2([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum_2([]int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}))
	fmt.Println(threeSum_2([]int{}))
	fmt.Println()
	fmt.Println(threeSum_3([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum_3([]int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}))
	fmt.Println(threeSum_3([]int{}))
	fmt.Println(threeSum_3([]int{-1,0,1,2,-1,-4}))
	fmt.Println(threeSum_3([]int{0, 0, 0}))
}

// 使用两重for结合set实现
func threeSum(nums []int) [][]int {
	res := [][]int{}
	target := 0
	var mapTemp map[int]int
	mapRR := make(map[string]bool)
	var keyTemp string
	for i:=0; i<len(nums)-1; i++ {
		// 剩下的数组
		// [-1, 0, 1, 2, -1, -4]
		target = -nums[i]
		// 用一层循环
		mapTemp = make(map[int]int)
		for j := i+1; j<len(nums); j++ {
			if _, ok := mapTemp[nums[j]]; ok {
				// 判重
				keyTemp = getKey(nums[i],mapTemp[nums[j]], nums[j])
				if ok := mapRR[keyTemp]; ok {
					continue
				}
				mapRR[keyTemp] = true
				res = append(res, []int{nums[i], mapTemp[nums[j]], nums[j]})
				continue
			}
			mapTemp[target - nums[j]] = nums[j]
		}

	}
	return res
}

// 如何高效去重
func getKey(a int, b int, c int) string {
	temp := ""
	if a > b {
		if b > c {
			return strconv.Itoa(c) + "-" + strconv.Itoa(b) + "-" + strconv.Itoa(a)
		} else {
			if a > c {
				// b < c < a
				return strconv.Itoa(b) + "-" + strconv.Itoa(c) + "-" + strconv.Itoa(a)
			} else {
				// b < a < c
				return strconv.Itoa(b) + "-" + strconv.Itoa(a) + "-" + strconv.Itoa(c)
			}
		}
	} else {
		// a <= b
		if c < a {
			// c < a < b
			return strconv.Itoa(c) + "-" + strconv.Itoa(a) + "-" + strconv.Itoa(b)
		} else {
			if b < c {
				// a < b < c
				return strconv.Itoa(a) + "-" + strconv.Itoa(b) + "-" + strconv.Itoa(c)
			} else {
				// a < c < b
				return strconv.Itoa(a) + "-" + strconv.Itoa(c) + "-" + strconv.Itoa(b)
			}
		}
	}
	return temp
}

// 先排序，然后find
func threeSum_2(nums []int) [][]int {
	res := [][]int{}
	// 先sort
	sortedNums := QSort2(nums, 0, len(nums) - 1)
	// 三元组是有序的，三元组第一个值作为key，第二个值作为value
	mapRR := make(map[string]int, 0)
	var x, y int
	for i:=0; i<len(sortedNums)-1; i++ {
		x, y = i+1, len(sortedNums) - 1
		// 两边向中间夹
		for x < y {
			if sortedNums[i] + sortedNums[x] + sortedNums[y] == 0 {
				key := getKey2(sortedNums[i], sortedNums[x], sortedNums[y])
				if val, ok := mapRR[key]; ok && val == sortedNums[x] {
					// 重复
					x++; y--
					continue
				}
				mapRR[key] = sortedNums[x]
				res = append(res, []int{sortedNums[i], sortedNums[x], sortedNums[y]})
				x++; y--
				continue
			} else if sortedNums[i] + sortedNums[x] + sortedNums[y] > 0 {
				// 右边向左移
				y--; continue
			} else {
				x++; continue
			}
		}
	}
	return res
}

func getKey2(a int, b int, c int) string {
	return strconv.Itoa(a) + "-" + strconv.Itoa(b) + "-" + strconv.Itoa(c)
}

// sorted and find，不单独使用map来去重，使用逻辑来判断去重。时间复杂度O（N^2）,空间复杂度O（1）
// 判重逻辑
// 1. a[i] > 0 直接跳过i
// 2. a[i] == a[i-1]，直接跳过i
// 3. 当sum==0，nums[L] == nums[L+1], L++; nums[R] == nums[R-1], R--
func threeSum_3(nums []int) [][]int {
	res := [][]int{}
	sortedNums := QSort2(nums, 0, len(nums) - 1)
	L, R, sum := 0, 0, 0
	for i:=0; i<len(sortedNums); i++ {
		if sortedNums[i] > 0 || (i >= 1 && sortedNums[i] ==  sortedNums[i-1])  {
			// 满足逻辑1和逻辑2
			continue
		}
		L = i+1
		R = len(sortedNums) - 1
		for L < R {
			sum = sortedNums[i] + sortedNums[L] + sortedNums[R]
			if sum == 0 {
				// 逻辑3
				if L + 1 < R && sortedNums[L] == sortedNums[L+1] {
					L++; continue
				}
				if R - 1 > L && sortedNums[R] == sortedNums[R-1] {
					R--; continue
				}
				// 去重后的结果
				res = append(res, []int{sortedNums[i], sortedNums[L], sortedNums[R]})
				L++; R--
			} else if sum < 0 {
				L++
			} else {
				R--
			}
		}
	}
	return res
}