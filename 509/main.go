package main

import "fmt"

func main() {
	res := fib(2)
	fmt.Println(res)
}

func fib(N int) int {
	if N == 0 || N == 1 {
		return N
	}
	// first 是第0个和第1个元素
	first := [][2]int{[2]int{0, 1}, [2]int{0, 0}}
	// temp为系数
	temp := [][2]int{[2]int{0, 1}, [2]int{1, 1}}
	res := matrix22_pow(temp, N-1)
	return matrix22_mul(first, res)[0][1]
}

func matrix22_pow(x [][2]int, n int) [][2]int {
	r := x
	res := [][2]int{[2]int{1, 0}, [2]int{0, 1}}
	for n != 0 {
		if n&1 == 1 {
			// 最低二进制位为1
			res = matrix22_mul(res, r)
		}
		// 2维矩阵相乘
		r = matrix22_mul(r, r)
		n >>= 1
	}
	return res
}

func matrix22_mul(x, y [][2]int) [][2]int {
	temp := make([][2]int, 2)
	temp[0][0] = x[0][0]*y[0][0] + x[0][1]*y[0][1]
	temp[0][1] = x[0][0]*y[0][1] + x[0][1]*y[1][1]
	temp[1][0] = x[1][0]*y[0][0] + x[1][1]*y[0][1]
	temp[1][1] = x[1][0]*y[1][0] + x[1][1]*y[1][1]
	return temp
}
