package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(10))
	fmt.Println(isPowerOfTwo(218))
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(8))
	fmt.Println(isPowerOfTwo(1024))
}

//func isPowerOfTwo(n int) bool {
//	if n <= 0 {
//		return false
//	}
//	x := n & (n - 1)
//	if x > 0 {
//		return false
//	}
//	return true
//}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
