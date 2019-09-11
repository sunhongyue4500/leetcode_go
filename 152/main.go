package main

import (
	"fmt"
	"log"
)

func main() {
	input := []int{
		-2, 3, -4,
	}

	log.Println(maxProduct(input))

}

func maxProduct(nums []int) int {
	dp := make([][2]int, len(nums))
	dp[0] = [2]int{nums[0], nums[0]}
	res := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] >= 0 {
			dp[i+1][0] = max(dp[i][0]*nums[i+1], nums[i+1])
			dp[i+1][1] = min(dp[i][1]*nums[i+1], nums[i+1])
		} else {
			dp[i+1][0] = max(dp[i][1]*nums[i+1], nums[i+1])
			dp[i+1][1] = min(dp[i][0]*nums[i+1], nums[i+1])
		}
		if res < dp[i+1][0] {
			res = dp[i+1][0]
		}
	}
	fmt.Println(dp)
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
