package main

import (
	"fmt"
	"reflect"
)

func main() {
	iteration1("abc")
	iteration2("abc")
	fmt.Println(reverse("abcdef"))
}

/**************string 遍历*****************/
func iteration1(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i], reflect.TypeOf(str[i])) // 这样访问是byte
	}
}

func iteration2(str string) {
	for _, str := range str {
		fmt.Println(str, reflect.TypeOf(str)) // 这样访问是rune
	}
}

/**************string 原地反转*****************/
func reverse(str string) string {
	i := 0
	j := len(str) - 1
	runeArray := []rune(str)
	for i < j {
		temp := runeArray[i]
		runeArray[i] = runeArray[j]
		i++
		runeArray[j] = temp
		j--
	}
	return string(runeArray)
}
