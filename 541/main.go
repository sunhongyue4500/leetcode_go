package main

import "fmt"

func main() {

	fmt.Println(reverseStr("abcdefg", 2))
	fmt.Println(reverseStr("abcdefg", 1))

	fmt.Println(reverseStr("krmyfshbspcgtesxnnljhfursyissjnsocgdhgfxubewllxzqhpasguvlrxtkgatzfybprfmmfithphckksnvjkcvnsqgsgosfxc", 20))
	fmt.Println()
	fmt.Println(reverseStr2("abcdefg", 2))
	fmt.Println(reverseStr2("abcdefg", 1))
	fmt.Println(reverseStr2("krmyfshbspcgtesxnnljhfursyissjnsocgdhgfxubewllxzqhpasguvlrxtkgatzfybprfmmfithphckksnvjkcvnsqgsgosfxc", 20))
}

/*
给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。
如果剩余少于 k 个字符，则将剩余的所有全部反转。
如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样。
**** **** **  k=2
直接写，效率很低
*/
func reverseStr(s string, k int) string {
	if k == 1 {
		return s
	}
	n := len(s)
	startIdx := 0
	res := ""
	for startIdx < n {
		if startIdx+2*k < n {
			// 前k个反转
			res = res + reverse(s[startIdx:startIdx+k]) + s[startIdx+k:startIdx+2*k]
			startIdx = startIdx + 2*k
		} else if startIdx+k >= n {
			// 全部反转
			res = res + reverse(s[startIdx:])
			break
		} else {
			// 只反转前k个
			res = res + reverse(s[startIdx:startIdx+k]) + s[startIdx+k:]
			break
		}
	}
	return res
}

func reverse(str string) string {
	res := ""
	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}
	return res
}

// 以2k为单位，反转前k个，如果不足2k，且大于k，反转前k个，如果不足k个，反转剩余的
// abcdef，从两边向中间夹，实现反转
func reverseStr2(s string, k int) string {
	runeArr := []rune(s)
	for start := 0; start < len(runeArr); start += 2 * k {
		i := start
		j := i + k - 1
		if j > len(s)-1 {
			j = len(s) - 1
		}
		for i < j { // 反转
			temp := runeArr[i]
			runeArr[i], i = runeArr[j], i+1
			runeArr[j], j = temp, j-1
		}
	}
	return string(runeArr)
}
