package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("345"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("9"))

	fmt.Println(letterCombinations_2("23"))
	fmt.Println(letterCombinations_2("345"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	dic := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := []string{}
	helper(digits, &res, 0, "", dic)

	return res
}

func helper(digits string, res *[]string, level int, curStr string, dic map[byte]string) {
	// terminator
	if level >= len(digits) {
		*res = append(*res, curStr)
		return
	}

	// process
	temp := digits[level]
	tempStr := dic[temp]
	for _, tempR := range tempStr {
		// drill down
		helper(digits, res, level+1, curStr+string(tempR), dic)
	}
	// clear status
}

type Queue []string

func (queue *Queue) add(str string) {
	*queue = append(*queue, str)
}

func (queue *Queue) isEmpty() bool {
	return len(*queue) == 0
}

func (queue *Queue) peek() string {
	if queue.isEmpty() {
		return ""
	}
	return (*queue)[0]
}

func (queue *Queue) remove() string {
	if queue.isEmpty() {
		return ""
	}
	temp := (*queue)[0]
	*queue = (*queue)[1:]
	return temp
}

func letterCombinations_2(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	dic := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := Queue{}
	res.add("")
	for i := 0; i < len(digits); i++ {
		aa := dic[digits[i]]
		size := len(res)
		for k := 0; k < size; k++ {
			t := res.remove()
			for j := 0; j < len(aa); j++ {
				res.add(t + string(aa[j]))
			}
		}
	}
	return res
}
