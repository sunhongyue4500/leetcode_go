package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(4294967295))
	fmt.Println(hammingWeight(0b11111111111111111111111111111111))
	fmt.Println(0b11111111111111111111111111111111)
}

func hammingWeight(num uint32) int {
	x, ctn := num, 0
	for x != 0 {
		x = x & (x - 1)
		ctn++
	}
	return ctn
}
