package main

import (
	"fmt"
	"strings"
)

// 1. 暴力 N^3
// 2. 滑动窗口
func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

// pwwkew
// abcabcbb
func lengthOfLongestSubstring(s string) int {
	map1 := make(map[byte]int)
	start := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if val, ok := map1[s[i]]; ok { // 遇到重复，从前面重复的下一个开始
			if val+1 > start {
				start = val + 1
			}
		}
		map1[s[i]] = i
		if res < i-start+1 { // 记录每次是否有最大值
			res = i - start + 1
		}
	}
	return res
}

func lengthOfLongestSubstring2(s string) int {
	// 定义游标尺寸大小,游标的左边位置
	window, start := 0, 0

	// 循环字符串
	for key := 0; key < len(s); key++ {
		// 查看当前字符串是否在游标内
		isExist := strings.Index(string(s[start:key]), string(s[key]))

		// 如果不存在游标内部,游标长度重新计算并赋值
		if isExist == -1 {
			if key-start+1 > window {
				window = key - start + 1
			}
		} else { //存在，游标开始位置更换为重复字符串位置的下一个位置
			start = start + 1 + isExist
		}
	}
	return window
}
