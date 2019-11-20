package main

import (
	"fmt"
)

type mm struct {
	a bool
	b float32
}

func main() {
	test := Constructor()
	test.Insert("apple")
	fmt.Println(test.Search("apple"))
	fmt.Println(test.Search("app"))
	fmt.Println(test.StartsWith("app"))
	test.Insert("app")
	fmt.Println(test.Search("app"))
	//fmt.Println(test.StartsWith("ap"))
	//fmt.Println(test.StartsWith("tp"))
	//test.Insert("golang")
	//fmt.Println(test.Search("golang"))
}

// 用数组实现的字典树，比较浪费空间，查询效率较高
type Trie struct {
	child [26]*Trie
	val   byte
	end   bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	var j byte
	for i := 0; i < len(word); i++ {
		j = word[i] - 'a'
		if cur.child[j] == nil {
			cur.child[j] = &Trie{val: word[i]}
		}
		cur = cur.child[j]
	}
	// 需要标记这是一个词
	cur.end = true
}

/** Returns if the word is in the trie. 必须完全匹配*/
func (this *Trie) Search(word string) bool {
	cur := this
	var j byte
	for i := 0; i < len(word); i++ {
		j = word[i] - 'a'
		if cur.child[j] == nil {
			return false
		}
		cur = cur.child[j]
	}
	if !cur.end {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	var j byte
	for i := 0; i < len(prefix); i++ {
		j = prefix[i] - 'a'
		if cur.child[j] == nil {
			return false
		}
		cur = cur.child[j]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
