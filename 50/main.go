package main

import "fmt"

func main() {
	// test case
	//fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow_2(2.10000, 3))
	fmt.Println(myPow_2(2.00000, -2))
	fmt.Println(myPow_2(2.00000, 0))

}

//func myPow(x float64, n int) float64 {
//	// 递归终止条件
//	if n == 0 {
//		return 1
//	}
//	if n < 0 {
//		return 1.0 / myPow(x, -n)
//	}
//
//	// drill down
//	r := myPow(x, n / 2)
//
//	if n & 1 == 1 {
//		// 奇数
//		return r * r * x
//	}
//	return r * r
//}

func myPow(x float64, n int) float64 {
	// 看n的二进制位
	res, temp := 1.0, x
	m := n
	if n < 0 {
		m = -n
	}
	for m != 0 {
		if m&1 == 1 {
			// 奇数
			res *= temp
		}
		m >>= 1
		temp *= temp
	}
	if n < 0 {
		return 1.0 / res
	}
	return res
}

/* 二刷 递归解法 logN*/
//func myPow_2(x float64, n int) float64 {
//	if n == 0 {
//		return 1.0
//	}
//	// 先考虑正数
//	if n > 0 {
//		temp := myPow(x, n / 2)
//		if n & 1 == 0 {
//			// 偶数
//			return temp * temp
//		} else {
//			// 奇数
//			return temp * temp * x
//		}
//	}
//	// 负数
//	return 1.0 / myPow(x, -n)
//}

/* 二刷 迭代解法 logN*/
func myPow_2(x float64, n int) float64 {
	res, temp := 1.0, x
	m := n
	if n < 0 {
		m = -n
	}
	for m != 0 {
		if m&1 == 1 {
			res *= temp
		}
		m = m >> 1
		temp *= temp
	}
	if n < 0 {
		return 1.0 / res
	}
	return res
}
