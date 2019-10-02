package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("1中国abc", "abc国中1"))
}

func isAnagram(s string, t string) bool {
	mapS := make(map[string]int, 0)
	mapT := make(map[string]int, 0)
	if len(s) != len(t) {return false}
	var tempS, tempT string
	for i:=0; i<len(s); i++ {
		tempS = string(s[i])
		tempT = string(t[i])
		mapS[tempS]++
		mapT[tempT]++
	}
	if len(mapS) != len(mapT) {return false}
	for key, tempS := range mapS {
		if tempT, ok := mapT[key]; ok && tempT == tempS{
			continue
		}
		return false
	}
	return true
}