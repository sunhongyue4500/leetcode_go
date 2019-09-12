package main

import "fmt"

func main() {
	fmt.Println(isValid("{{}}"))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid(""))
	fmt.Println(isValid("]"))
	fmt.Println(isValid("("))
}

// 使用栈
func isValid(s string) bool {
	stack := make([]byte, 0)
	symbolMap := map[byte]byte{'}' : '{', ']' : '[', ')' : '('}
	var val byte; var ok bool
	for i:=0; i<len(s); i++ {
		if val, ok = symbolMap[s[i]]; ok {
			// 判断匹配
			if len(stack) == 0 || stack[len(stack) - 1] != val {
				return false
			}
			// pop
			stack = stack[0:len(stack)-1]
		} else {
			// push
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
