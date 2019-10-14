package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 0))
	fmt.Println(fourSum([]int{1,-2,-5,-4,-3,3,3,5}, -11))
	fmt.Println(fourSum([]int{-3,-2,-1,0,0,1,2,3}, 0))
}


// 和3sum的判重还是有区别的
// 双指针法，头尾向中间夹
// 去重
// 1. nums[i] > target || nums[i] + nums[j] > target (这个不成立，比如-5 -4 -3 -2这种)
// 1.
// 		i+3 < n && nums[i] + nums[i+1] + nums[i+2] + nums[i+3] > target; break
// 		n-3 > i && nums[i] + nums[n-1] + nums[n-2] + nums[n-3] < target; continue
//		j+2 < n && nums[i] + nums[j] + nums[j+1] + nums[j+2] > target; break
//		nums[i] + nums[j] + nums[n-1] + nums[n-2] < target; continue
// 2. nums[i] == nums[i-1], i++
// 3. nums[j] == nums[j-1], j++
// 4. sum == target && nums[L] == nums[L+1], L++
// 5. sum == target && nums[R] == nums[R-1], R--
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	// 先排序
	arr := nums
	sort.Ints(arr)
	L, R, sum, n := 0, 0, 0, len(arr)
	if n < 4 {return res}	// 不足四个元素
	// 两层for
	for i:=0; i<len(arr)-3; i++ {
		// 剪枝
		if i+3 < len(arr) && nums[i] + nums[i+1] + nums[i+2] + nums[i+3] > target {
			// 最小的都比target大，break
			break
		}
		if n-3 > i && nums[i] + nums[n-1] + nums[n-2] + nums[n-3] < target {
			// 加上最大的也小于target，continue
			continue
		}
		// 去重
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		for j:=i+1; j<len(arr)-2; j++ {
			// 剪枝
			if j + 2 < n && nums[i] + nums[j] + nums[j+1] + nums[j+2] > target {
				break
			}
			if n-2 > j && nums[i] + nums[j] + nums[n-1] + nums[n-2] < target {
				continue
			}
			// 去重
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}
			L = j+1
			R = len(arr) - 1
			for L < R {
				sum = arr[i] + arr[j] + arr[L] + arr[R]
				if sum == target {
					if L+1 < R && arr[L] == arr[L+1] {
						L++; continue
					}
					if R-1 > L && arr[R] == arr[R-1] {
						R--; continue
					}
					res = append(res, []int{arr[i], arr[j], arr[L], arr[R]})
					L++; R--
				} else if sum > target {
					R--
				} else {
					L++
				}
			}
		}
	}
	return res
}