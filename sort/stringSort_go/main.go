package main

import (
	"fmt"
	"sort"
)

type runeSlice []rune

func (rs runeSlice) Len() int {
	return len(rs)
}

func (rs runeSlice) Less(i, j int) bool {
	return rs[i] < rs[j]
}

func (rs runeSlice) Swap(i, j int) {
	temp := rs[i]
	rs[i] = rs[j]
	rs[j] = temp
}

func main() {
	var temp runeSlice
	str := "bcazdf"
	temp = []rune(str)
	fmt.Println(temp)
	sort.Sort(temp)
	fmt.Println(temp)
}
