package test

import "testing"

func Benchmark367_2Fen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPerfectSquare2Fen(100)
	}
}

func Benchmark367_Newton(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPerfectSquareNewTon(100)
	}
}

func isPerfectSquare2Fen(num int) bool {
	mid, left, right := 0, 0, num
	for left <= right {
		mid = left + (right-left)/2
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// 牛顿迭代法
func isPerfectSquareNewTon(num int) bool {
	r := num
	for r*r > num {
		r = (r*r + num) / (2 * r)
	}
	if r*r == num {
		return true
	}
	return false
}
