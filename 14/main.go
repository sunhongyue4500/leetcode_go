package main

import (
	"fmt"
	"sort"
	"strings"
)

type Trie struct {
	Nodes [26]*Trie
	total int // 记录这一层有多少个子节点
	val   byte
	isEnd bool
}

func (trie *Trie) insert(str string) {
	temp := trie
	for i := 0; i < len(str); i++ {
		if temp.Nodes[str[i]-'a'] == nil {
			temp.Nodes[str[i]-'a'] = &Trie{
				Nodes: [26]*Trie{},
			}
			temp.val = str[i]
			temp.total++
		}
		temp = temp.Nodes[str[i]-'a'] // 下探一层
	}
	temp.isEnd = true
}

func cons(strs []string) *Trie {
	trie := Trie{
		Nodes: [26]*Trie{},
	}
	for i := 0; i < len(strs); i++ {
		trie.insert(strs[i])
	}
	return &trie
}

func (trie *Trie) search(str string) bool {
	temp := trie
	for i := 0; i < len(str); i++ {
		if temp.Nodes[str[i]-'a'] == nil {
			return false
		}
		temp = temp.Nodes[str[i]-'a'] // 下探一层
	}
	return temp.isEnd
}

func (trie *Trie) startWith(str string) bool {
	temp := trie
	for i := 0; i < len(str); i++ {
		if temp.Nodes[str[i]-'a'] == nil {
			return false
		}
		temp = temp.Nodes[str[i]-'a'] // 下探一层
	}
	return !temp.isEnd
}

func main() {
	str := []string{"flower", "flow", "fldeft", "floht", "fli", "flabac"}
	str2 := []string{"", "floht"}
	str3 := []string{"aa", "a"}
	str4 := []string{"a"}
	str5 := []string{""}
	fmt.Println(longestCommonPrefix(str))
	fmt.Println(longestCommonPrefix(str2))
	fmt.Println(longestCommonPrefix2(str))
	fmt.Println(longestCommonPrefix2(str2))
	//fmt.Println(prefix("flow", "floht"))

	trie := cons(str)
	fmt.Println(trie.search("flo"))
	fmt.Println(trie.startWith("flo"))
	trie.insert("flo")
	fmt.Println(trie.search("flo"))
	fmt.Println(longestCommonPrefix3(str))
	fmt.Println(longestCommonPrefix3(str2))
	fmt.Println(longestCommonPrefix3(str3))
	// sort and compare
	fmt.Println("sort and compare")
	fmt.Println(longestCommonPrefix4(str))
	fmt.Println(longestCommonPrefix4(str2))
	fmt.Println(longestCommonPrefix4(str3))
	fmt.Println(longestCommonPrefix4(str4))
	fmt.Println(longestCommonPrefix4(str5))
}

// 1. 暴力，直接从头往后, O(MN)，M为最长公共前缀，N为字符串个数
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pos := 0
	var temp byte
	count := len(strs[0])
loop:
	for pos < count {
		temp = strs[0][pos]
		for j := 1; j < len(strs); j++ {
			if pos > len(strs[j])-1 || strs[j][pos] != temp {
				break loop
			}
		}
		pos++
	}
	return strs[0][:pos]
}

// 2. 水平扫描法，第一个依次和后面的所有单词取公共前缀
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		//res = prefix(res, strs[i])
		// 关键代码
		for strings.Index(strs[i], res) != 0 { // 完全匹配
			res = res[:len(res)-1] // 如果不完全匹配，就缩短
			if res == "" {
				return ""
			}
		}
	}
	return res
}

func prefix(str1 string, str2 string) string {
	cnt := 0
	for cnt < len(str1) {
		if cnt > len(str2)-1 || str1[cnt] != str2[cnt] {
			break
		}
		cnt++
	}
	return str1[:cnt]
}

// Trie
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	trie := Trie{
		Nodes: [26]*Trie{},
	}
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return ""
		}
		trie.insert(strs[i])
	}

	temp := &trie
	prefix := ""
	for !temp.isEnd && temp.total == 1 { // 以最短的为准，最长前缀不回超过最短的长度
		prefix += string(temp.val)
		temp = temp.Nodes[temp.val-'a']
	}
	return prefix
}

// sort and compare
func longestCommonPrefix4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	temp := sort.StringSlice(strs)
	temp.Sort()
	// 比较第一个和最后一个
	strA := temp[0]
	strB := temp[len(temp)-1]
	length := len(strA)
	if length > len(strB) {
		length = len(strB)
	}
	res := ""
	for i := 0; i < length; i++ {
		if strA[i] != strB[i] {
			break
		}
		res += string(strA[i])
	}
	return res
}
