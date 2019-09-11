package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare1(10000))
	fmt.Println(isPerfectSquare2(10000))
}

func isPerfectSquare1(num int) bool {
	mid, left, right := 0, 0, num
	count := 0
	for left <= right {
		count++
		mid = left + (right-left)/2
		if mid*mid == num {
			fmt.Println("1count：", count)
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("1count：", count)
	return false
}

// 牛顿迭代法
func isPerfectSquare2(num int) bool {
	r := num
	count := 0
	for r*r > num {
		count++
		r = (r*r + num) / (2 * r)
	}
	fmt.Println("2count：", count)
	if r*r == num {
		return true
	}
	return false
}
