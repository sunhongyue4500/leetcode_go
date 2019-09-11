package main

import (
	"fmt"
	"testing"
)

type Stack struct {
	stackSlice []byte
}

func (stack *Stack) Push(ele byte) {
	stack.stackSlice = append(stack.stackSlice, ele)
}

func (stack *Stack) Equal(stack2 *Stack) bool {
	if len(stack.stackSlice) != len(stack2.stackSlice) {
		return false
	}
	for idx, val := range stack.stackSlice {
		if val != stack2.stackSlice[idx] {
			return false
		}
	}
	return true
}

func (stack *Stack) Pop() (ele byte) {
	if len(stack.stackSlice) == 0 {
		return '0'
	}
	temp := stack.stackSlice[len(stack.stackSlice)-1]
	// reslice
	stack.stackSlice = stack.stackSlice[0 : len(stack.stackSlice)-1]
	return temp
}

func backspaceCompare(S string, T string) bool {
	myStack1 := Stack{stackSlice: make([]byte, 0)}
	myStack2 := Stack{stackSlice: make([]byte, 0)}
	lenS := len(S)
	lenT := len(T)
	for i := 0; i < lenS; i++ {
		if S[i] == byte('#') {
			myStack1.Pop()
			continue
		}
		myStack1.Push(S[i])
	}
	for i := 0; i < lenT; i++ {
		if T[i] == byte('#') {
			myStack2.Pop()
			continue
		}
		myStack2.Push(T[i])
	}
	fmt.Println(myStack1.stackSlice)
	fmt.Println(myStack2.stackSlice)
	return myStack1.Equal(&myStack2)
}

func TestStack(t *testing.T) {
	// myStack := Stack{stackSlice: make([]byte, 0)}
	// myStack.Push('a')
	// myStack.Push('b')
	// myStack.Push('c')
	// myStack.Pop()
	// t.Log(myStack)

	// slice1 := make([]int, 0, 10)
	// slice2 := make([]int, 0, 11)
	// t.Log("equal:", slice1 == slice2)

	S := "ab##"
	T := "c#d#"
	// S = "ab#c"
	// T = "ad#c"
	// S = "a#c"
	// T = "b"
	// S = "#####"
	// T = "a########"
	S = ""
	T = ""
	S = "xywrrmp"
	T = "xywrrmu#p"
	// S = "bbbextm"
	// T = "bbb#extm"
	t.Log(backspaceCompare2(S, T))
}

func TestReverseStringWithSharp(t *testing.T) {
	S := "abc#d#"
	// S = "a##c"
	// S = "c#d#"
	// S = "#a#c"
	S = "abcd"
	S = "###abc#####"
	i := len(S) - 1

	skipI := 0

	// 从后往前遍历
	for i >= 0 {
		// 找到s，第一个不是#的字符
		for i >= 0 {
			if S[i] == byte('#') {
				skipI++
				i--
			} else if skipI > 0 {
				i--
				skipI--
			} else {
				skipI = 0
				break
			}
		}
		if i >= 0 {
			t.Log("找到result:", S[i])
		}
		i--
	}
}

func backspaceCompare2(S string, T string) bool {
	i := len(S) - 1
	j := len(T) - 1
	skipI, skipJ := 0, 0

	// 从后往前遍历,|| 关系确保两个都一个<0，一个>=0
	for i >= 0 || j >= 0 { // 全都遍历完
		// 找到S第一个不是#的字符
		for i >= 0 {
			if S[i] == byte('#') {
				i--
				skipI++
			} else if skipI > 0 {
				i--
				skipI--
			} else {
				break
			}
		}
		// 找到T第一个不是#的字符
		for j >= 0 {
			if T[j] == byte('#') {
				j--
				skipJ++
			} else if skipJ > 0 {
				j--
				skipJ--
			} else {
				break
			}
		}

		if i >= 0 {
			fmt.Println("找到resultS:", S[i], i)
		}
		if j >= 0 {
			fmt.Println("找到resultT:", T[j], j)
		}
		if i >= 0 && j >= 0 && S[i] != T[j] {
			return false
		}
		// 如果i，j其中一个首先遍历完成，正负相反
		if (i >= 0) != (j >= 0) {
			return false
		}
		i--
		j--
	}
	return true
}
