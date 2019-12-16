package main

import (
	"fmt"
	"math"
)

func main() {
	test1 := []int{3, 6, 7}
	test2 := []int{9, 9, 9}
	test3 := []int{0}
	test4 := []int{3, 9, 9}
	test5 := []int{3, 4, 9}
	test6 := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	fmt.Println(plusOne(test1))
	fmt.Println(plusOne(test2))
	fmt.Println(plusOne(test3))
	fmt.Println(plusOne(test4))
	fmt.Println(plusOne(test5))
	fmt.Println(plusOne(test6), len(test6), math.MaxInt64)

	//fmt.Println(pow(10, 4))
	//fmt.Println(pow(10, 0))
	//fmt.Println(pow(2, 10))
	//fmt.Println(pow(2, 9))
	//fmt.Println(pow(10, 5))
}

// 1. 计算加1，会溢出
// 两遍for，计算数值pow(a, n)，这个是lgN的时间复杂度
//func plusOne(digits []int) []int {
//	//temp := 0
//	counter := 0
//	sum := 0
//
//	for i := len(digits) - 1; i >= 0; i-- {
//		sum += digits[i] * pow(10, counter)
//		counter++
//	}
//	sum += 1
//	// 将数字转换成数组
//	res := make([]int, len(digits)+1)
//	for i := len(res) - 1; i >= 0 && sum > 0; i-- {
//		temp1 := sum % 10
//		sum = sum / 10
//		res[i] = temp1
//	}
//	// 去除开头的0
//	if res[0] == 0 && len(res) > 1 {
//		res = res[1:]
//	}
//	return res
//}
//
//// logN
//func pow(a, n int) int {
//	current := a
//	res := 1
//	for i := n; i >= 1; i >>= 1 {
//		if i&1 == 1 {
//			res *= current
//		}
//		current = current * current
//	}
//	return res
//}

// 2. 直接在数组上操作
//执行用时 :0 ms, 在所有 golang 提交中击败了100.00%的用户
//内存消耗 :2.1 MB, 在所有 golang 提交中击败了99.56%的用户
func plusOne(digits []int) []int {
	res := make([]int, len(digits)+1)
	length := len(digits)
	jinwei := 0
	for i := length - 1; i >= 0; i-- {
		if i == length-1 {
			jinwei = (digits[i] + 1) / 10
			res[i+1] = (digits[i] + 1) % 10
			continue
		}
		res[i+1] = (digits[i] + jinwei) % 10
		jinwei = (digits[i] + jinwei) / 10
	}
	if jinwei > 0 {
		res[0] = jinwei
	} else {
		// 去掉开头的0
		res = res[1:]
	}
	return res
}
