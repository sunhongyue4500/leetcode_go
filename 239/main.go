package main

import "fmt"

// 用slice实现
type Deque struct {
	data []int
}

func (deque *Deque)AddFirst(ele int)  {
	deque.data = append([]int{ele}, deque.data...)
}

func (deque *Deque)AddLast(ele int) {
	deque.data = append(deque.data, ele)
}

func (deque *Deque)RemoveFirst() int {
	res := deque.data[0]
	deque.data = deque.data[1:]
	return res
}

func (deque *Deque)RemoveLast() int {
	res := deque.data[len(deque.data) - 1]
	deque.data = deque.data[:len(deque.data) - 1]
	return res
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	// window存储的下标
	window := Deque{[]int{}}
	res := []int{}
	for i:=0; i<len(nums); i++ {
		// window 存储的是索引
		if i >= k && window.data[0] <= i - k {
			window.RemoveFirst()
		}
		for len(window.data) > 0 && nums[window.data[len(window.data) - 1]] <= nums[i] {
			window.RemoveLast()
		}
		window.AddLast(i)
		if i >= k - 1 {
			// 往结果中添加
			res = append(res, nums[window.data[0]])
		}
	}
	return res
}

// 用切片模拟的双端队列
func maxSlidingWindow_2(nums []int, k int) []int {
	// 1,3,-1,-3,5,3,6,7
	res := []int{}
	if len(nums) == 0 || k < 1 {
		return res
	}
	window := []int{} // 滑动窗口，存储的索引
	for i:=0; i<len(nums); i++ {
		if len(window) > 0 && i - k >= window[0] {
			// 需要向右滑动
			window = window[1:]
		}
		for len(window) > 0 {
			// 从后往前看，如果小就删除
			if nums[i] >= nums[window[len(window) - 1]] {
				// 移除最后一个
				window = window[:len(window)-1]
			} else {
				break
			}
		}
		window = append(window, i)
		if i + 1 >= k {
			// 将结果添加到res
			res = append(res, nums[window[0]])
		}
	}

	return res
}

// 解法二 双端队列 Deque
func maxSlidingWindow_3(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	window := make([]int, 0, k) // store the index of nums
	result := make([]int, 0, len(nums)-k+1)
	for i, v := range nums { // if the left-most index is out of window, remove it
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < v { // maintain window
			window = window[:len(window)-1]
		}
		window = append(window, i) // store the index of nums
		if i >= k-1 {
			result = append(result, nums[window[0]]) // the left-most is the index of max value in nums
		}
	}
	return result
}

func main() {
	// 1,3,-1,-3,5,3,6,7
	// [1] 3, [1, 3] -> [3]
	// [3] -1 -> [3, -1]
	// [3, -1] -3, 长度等于k，先把左侧的3踢出，剩下[-1]，从后向前遍历res和-3比较，比它小踢出，比他大停止，并把-3加入

	//test := []int{1,3,-1,-3,5,3,6,7}
	test := []int{1,3,1,2,0,5}
	res := maxSlidingWindow_2(test, 3)
	fmt.Println(res)
}