package main

import "fmt"

func main() {
	//fmt.Println(mySqrt2(9, 1e-3))
	//fmt.Println(mySqrt(1))
	fmt.Println(mySqrt3(0))
}

// 浮点版
func mySqrt2(x int, pres float64) float64 {
	mid, left, right := 0.0, 0.0, float64(x)
	for left <= right {
		mid = left + (right-left)/2
		// 1e-3表示10^-3
		if fabs(mid*mid, float64(x)) < pres {
			return mid
		} else if mid*mid > float64(x) {
			right = mid
		} else {
			left = mid
		}
		fmt.Printf("%f %f %f\n", left, right, mid)
	}
	return mid
}

func fabs(a float64, b float64) float64 {
	if a > b {
		return a - b
	}
	return b - a
}

func mySqrt(x int) int {
	mid, left, right := 1, 1, x
	res := 0
	for left <= right {
		mid = left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
			// 记录上次的mid
			res = mid
		}
		fmt.Printf("%d %d %d\n", left, right, mid)
	}
	fmt.Printf("%d %d\n", left, right)
	return res
}

// 牛顿迭代法
func mySqrt3(x int) int {
	r := x
	for r*r > x {
		r = (r*r + x) / (2 * r)
	}
	return r
}
