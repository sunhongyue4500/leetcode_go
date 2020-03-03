package main

import "fmt"

func main() {
	fmt.Println(reversePairs2([]int{1, 3, 2, 3, 1}))
	fmt.Println(reversePairs2([]int{2, 4, 3, 5, 1}))
	fmt.Println(reversePairs2([]int{1, 2, 7}))

	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
	fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
	fmt.Println(reversePairs([]int{1, 2, 7}))
}

func reversePairs(nums []int) int {
	if nums == nil || len(nums) < 2 {
		// 数组为空，或长度小于2，肯定不存在逆序对
		return 0 // 返回0
	}
	temp := make([]int, len(nums))
	return mergeSort(nums, 0, len(nums)-1, temp)
}

func mergeSort(nums []int, left, right int, temp []int) int {
	// terminator
	if left >= right {
		return 0
	}
	// prepare data
	mid := left + (right-left)>>1
	// conquer sub prob
	cnt := mergeSort(nums, left, mid, temp)
	cnt += mergeSort(nums, mid+1, right, temp)
	cnt += count(nums, left, mid, right)
	//for i := left; i <= mid; i++ {
	//	j := mid + 1
	//	for j <= right && float32(nums[i])/2 > float32(nums[j]) {
	//		j++
	//	}
	//	cnt += j - (mid + 1)
	//}
	// merge
	merge(nums, left, mid, right, temp)
	return cnt
}

// 合并索引范围内的数组
func merge(nums []int, start, mid, end int, temp []int) {
	// 下面这种写法每次merge都要建立数组，比较耗时
	//temp := make([]int, end-start+1) // 只需要这段长度
	i, j, k := start, mid+1, 0
	for i <= mid && j <= end {
		if nums[i] > nums[j] {
			temp[k], k, j = nums[j], k+1, j+1
		} else {
			temp[k], k, i = nums[i], k+1, i+1
		}
	}
	// 处理剩余的
	for i <= mid {
		temp[k], k, i = nums[i], k+1, i+1
	}
	for j <= end {
		temp[k], k, j = nums[j], k+1, j+1
	}
	// copy to nums
	//for p := 0; p < len(temp); p++ {
	//	nums[p+start] = temp[p]
	//}
	copy(nums[start:end+1], temp)
}

// 别人的解法
func reversePairs2(nums []int) int {
	if nums == nil || len(nums) < 2 {
		// 数组为空，或长度小于2，肯定不存在逆序对
		return 0 // 返回0
	}
	tmp := make([]int, len(nums))                  // 定义辅助数组tmp
	return sortAndCount(nums, 0, len(nums)-1, tmp) // 一边做归并排序，一边统计逆序对数量
}

// 辅助函数 归并两个有序子数组
// 输入数组arr，low到mid对应第一个有序子数组，mid+1到high对应第二个有序子数组，辅助数组tmp
func merge2(arr []int, low, mid, high int, tmp []int) {
	// 定义游标i，j，k分别指向两个有序子数组和辅助数组tmp
	i, j, k := low, mid+1, 0    // 初始化为各自的开始下标
	for i <= mid && j <= high { // 当满足条件时不断执行以下操作
		if arr[i] <= arr[j] { // 如果arr[i]小于arr[j]
			tmp[k] = arr[i] // 把i指向的数字放到tmp中
			i, k = i+1, k+1 // i和k都加一
		} else { // 否则说明arr[i]>arr[j]产生逆序对
			tmp[k] = arr[j] // 同时也需要把较小的数字加入到tmp[k]中
			k, j = k+1, j+1 // 且j和k都需要加1
		}
	}
	// 退出循环后，为了避免其中一个子数组还有未进行对比的数字，所以对两个子数组进行检查
	for i <= mid { // 把下标i指向的数字以及后面的数字依次加入到tmp中
		tmp[k] = arr[i]
		k, i = k+1, i+1
	}
	for j <= high { // 把子数组剩下的数字加入到tmp中
		tmp[k] = arr[j]
		k, j = k+1, j+1
	}
	// 这时tmp是合并后的有序数组,再拷贝回arr数组
	copy(arr[low:high+1], tmp)
}

// 辅助函数 统计重要逆序对
func count(arr []int, low, mid, high int) int {
	i, j := low, mid+1
	cnt := 0
	for i <= mid && j <= high {
		if arr[i] <= 2*arr[j] { // 可能会越界
			i++
		} else {
			cnt += mid - i + 1
			j++
		}
	}
	return cnt
}

// 辅助函数 同时进行归并排序和统计逆序对数量
func sortAndCount(arr []int, low, high int, tmp []int) int {
	cnt := 0        // 定义cnt统计逆序对数量
	if low < high { // 当low小于high时执行以下操作
		mid := low + (high-low)/2                  // 计算low和high中间下标mid
		cnt += sortAndCount(arr, low, mid, tmp)    // 递归调用自己处理low到mid之间的子数组，同时把返回的逆序对数量加到cnt上
		cnt += sortAndCount(arr, mid+1, high, tmp) // 同样的处理mid+1到high子数组,同时把返回的逆序对数量加到cnt上
		cnt += count(arr, low, mid, high)          // 把返回的逆序对数量加到cnt上
		merge2(arr, low, mid, high, tmp)           // 最后调用辅助函数去合并两个子数组
	}
	return cnt // 最后返回cnt即可
}
