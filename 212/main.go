package main

import (
	"fmt"
)

func main() {
	testBoart := [][]byte{
		[]byte{'o', 'a', 'a', 'n'},
		[]byte{'e', 't', 'a', 'e'},
		[]byte{'i', 'h', 'k', 'r'},
		[]byte{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	testBoart2 := [][]byte{
		[]byte{'a', 'a'},
	}

	words2 := []string{"a"} // 使用字典树要考虑去重

	testBoart3 := [][]byte{
		[]byte{'a', 'b'},
		[]byte{'a', 'a'},
	}

	words3 := []string{"aaab", "aaa", "aaba"} // 使用字典树要考虑去重

	fmt.Println(findAWord(testBoart, "ABFDE"))
	fmt.Println(findAWord(testBoart, "ABCCC"))
	fmt.Println(findWords(testBoart, words))

	fmt.Println(findWords2(testBoart, words))
	fmt.Println(findWords2(testBoart2, words2))
	fmt.Println(findWords2(testBoart3, words3))
}

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	for i := 0; i < len(words); i++ {
		// 查看每一个word是否在board中
		if findAWord(board, words[i]) {
			res = append(res, words[i])
		}
	}
	return res
}

func findAWord(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] {
				visited[i][j] = true
				if dfs(board, rows, cols, i, j, word, visited) {
					return true
				}
				// 恢复状态
				visited[i][j] = false
			}
		}
	}

	return false
}

func dfs(board [][]byte, row, col, i, j int, str string, visited [][]bool) bool {
	// terminator
	if len(str) == 1 {
		if board[i][j] == str[0] {
			return true
		}
		return false
	}
	// 如果长度不为1
	if board[i][j] != str[0] {
		return false
	}

	// 当前位置的字符是匹配的
	// process
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}

	for k := 0; k < 4; k++ {
		// 判断越界并且没有访问过
		if i+dx[k] >= 0 && i+dx[k] < row && j+dy[k] >= 0 && j+dy[k] < col && !visited[i+dx[k]][j+dy[k]] {
			visited[i+dx[k]][j+dy[k]] = true
			// drill down
			if dfs(board, row, col, i+dx[k], j+dy[k], str[1:], visited) {
				return true
			}
			// clear status
			visited[i+dx[k]][j+dy[k]] = false
		}
	}

	return false
}

/************************使用字典树*********************************/
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

func findWords2(board [][]byte, words []string) []string {
	res := []string{}
	trie := Constructor()
	for i := 0; i < len(words); i++ {
		// 构造字典树
		trie.Insert(words[i])
	}

	resMap := make(map[string]bool)

	row := len(board)
	col := len(board[0])
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			dfs2(board, row, col, i, j, string(board[i][j]), visited, &trie, &res, resMap)
		}
	}
	return res
}

func dfs2(board [][]byte, row, col, i, j int, stringb string, visited [][]bool, trie *Trie, res *[]string, resMap map[string]bool) {
	// terminator
	if !trie.StartsWith(stringb) {
		return
	}
	if trie.Search(stringb) {
		// 字典树包含这个word
		if _, ok := resMap[stringb]; !ok {
			// 去重
			*res = append(*res, stringb)
			fmt.Println(stringb)
			resMap[stringb] = true
		}
		//return // 不能return，还得往下继续搜索
	}

	// 当前位置的字符是匹配的
	// process
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}

	visited[i][j] = true
	for k := 0; k < 4; k++ {
		// 判断越界并且没有访问过
		if i+dx[k] >= 0 && i+dx[k] < row && j+dy[k] >= 0 && j+dy[k] < col && !visited[i+dx[k]][j+dy[k]] {
			// drill down
			// 四个方向都要尝试
			dfs2(board, row, col, i+dx[k], j+dy[k], stringb+string(board[i+dx[k]][j+dy[k]]), visited, trie, res, resMap)
		}
	}
	visited[i][j] = false
}

/************************使用字典树，每次下探一层*********************************/
