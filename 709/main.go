package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(byte('a'))
	fmt.Println(byte('z'))
	fmt.Println(byte('A'))
	fmt.Println(byte('Z'))
	fmt.Println(toLowerCase("Hello"))
	fmt.Println(toLowerCase("here"))
	fmt.Println(toLowerCase("Ldsaf333OVELY"))
	fmt.Println(toLowerCase(""))
	fmt.Println(toLowerCase2("Ldsaf333OVELY"))
	fmt.Println(toLowerCase2(""))
	fmt.Println(toLowerCase3("Ldsaf333OVELY"))
	fmt.Println(toLowerCase3(""))
}

func toLowerCase(str string) string {
	res := ""
	for i := 0; i < len(str); i++ {
		if str[i] < 65 || str[i] > 90 {
			// 非小写字母
			res = res + string(str[i])
			continue
		}
		res = res + string(str[i]+32)
	}
	return res
}

// rune数组
func toLowerCase2(str string) string {
	strRune := []rune(str)
	for idx, tempRune := range strRune {
		if tempRune >= 'A' && tempRune <= 'Z' {
			strRune[idx] = tempRune + 32
		}
	}
	return string(strRune)
}

// bytes.buffer
func toLowerCase3(str string) string {
	var buffer bytes.Buffer
	for _, tempRune := range str {
		if tempRune >= 'A' && tempRune <= 'Z' {
			buffer.WriteRune(tempRune + 32)
			continue
		}
		buffer.WriteRune(tempRune)
	}
	return buffer.String()
}
