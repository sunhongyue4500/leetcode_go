package main

import "fmt"

func main() {
	fmt.Println(countBits(10))
	fmt.Println(countBits2(10))
	fmt.Println(countBits3(10))
}

// 5
func countBits(num int) []int {
	res := make([]int, num+1)
	var temp, counter int
	for i := 0; i <= num; i++ {
		temp, counter = i, 0
		for temp != 0 {
			temp = temp & (temp - 1)
			counter++
		}
		res[i] = counter
	}
	return res
}

func countBits2(num int) []int {
	res := make([]int, num+1)
	var temp, counter int
	for i := 0; i <= num; i++ {
		temp, counter = i, 0
		for temp != 0 {
			if temp%2 != 0 {
				counter++
			}
			temp >>= 1
		}
		res[i] = counter
	}
	return res
}

func countBits3(num int) []int {
	res := make([]int, num+1)
	res[0] = 0
	for i := 1; i <= num; i++ {
		if i&1 == 1 {
			res[i] = res[1] + 1
		} else {
			res[i] = res[0] + 1
		}
	}
	return res
}
