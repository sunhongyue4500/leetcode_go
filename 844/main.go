package main

import "fmt"

func main() {
	fmt.Println(backspaceCompare_2("a#", ""))
	fmt.Println(backspaceCompare_2("ab#c", "ad#c"))
	fmt.Println(backspaceCompare_2("ab##", "c#d#"))
	fmt.Println(backspaceCompare_2("a##c", "#a#c"))
	fmt.Println(backspaceCompare_2("a#c", "b"))
	fmt.Println(backspaceCompare_2("bbb", "bbb#"))
}

//func backspaceCompare(S string, T string) bool {
//	stack1 := Stack{data: []byte{}}
//	stack2 := Stack{data: []byte{}}
//	for i := 0; i < len(S); i++ {
//		if S[i] != '#' {
//			stack1.Push(S[i])
//		} else {
//			stack1.Pop()
//		}
//	}
//	for i := 0; i < len(T); i++ {
//		if T[i] != '#' {
//			stack2.Push(T[i])
//		} else {
//			stack2.Pop()
//		}
//	}
//	return stack1.Equal(&stack2)
//}

// 使用O(1)空间复杂度，两个字符串从后往前遍历。如果遇到#，计数skip个数，如果不是#，看skip是否大于0
// 需不需要跳过，如果不是#且skip等于0，则比较两个字符串的这个元素是否一样，知道遍历完成。
func backspaceCompare_2(S string, T string) bool {
	i := len(S) - 1
	j := len(T) - 1
	skipI, skipJ := 0, 0
	// 注意是||，让两个都遍历完成
	for i >= 0 || j >= 0 {
		// 从S中找到第一个不是#，并且跳过需要跳过的元素
		for i >= 0 {
			if S[i] == '#' {
				skipI++
				i--
			} else if skipI > 0 {
				skipI--
				i--
			} else {
				break
			}
		}
		// 从T中找到第一个不是#，并且跳过需要跳过的元素
		for j >= 0 {
			if T[j] == '#' {
				skipJ++
				j--
			} else if skipJ > 0 {
				skipJ--
				j--
			} else {
				break
			}
		}
		// 一个先遍历完，另一个还有元素
		if (i >= 0) != (j <= 0) {
			return false
		}
		// 找到了S和T中的元素，进行判断
		if i >= 0 && j >= 0 && S[i] != T[j] {
			return false
		}
		i--
		j--
	}
	return true
}
