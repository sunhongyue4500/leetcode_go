package main

import "fmt"

func main() {
	str := "abc"
	fmt.Println(string(str[0]))

	testBoart := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	test2 := [][]byte{
		[]byte{'C', 'A', 'A'},
		[]byte{'A', 'A', 'A'},
		[]byte{'B', 'C', 'D'},
	}

	test3 := [][]byte{
		[]byte{'A', 'A'},
	}

	fmt.Println('A', 'a', testBoart)
	fmt.Println(exist(testBoart, "ABCCED"))
	fmt.Println(exist(testBoart, "SEE"))
	fmt.Println(exist(testBoart, "ABCB"))
	fmt.Println(exist(testBoart, "SNO"))
	fmt.Println(exist(test2, "AAB"))
	fmt.Println()
	fmt.Println(exist2(testBoart, "ABCCED"))
	fmt.Println(exist2(testBoart, "SEE"))
	fmt.Println(exist2(testBoart, "ABCB"))
	fmt.Println(exist2(testBoart, "SNO"))
	fmt.Println(exist2(test2, "AAB"))
	fmt.Println()
	fmt.Println(exist3(testBoart, "ABCCED"))
	fmt.Println(exist3(testBoart, "SEE"))
	fmt.Println(exist3(testBoart, "ABCB"))
	fmt.Println(exist3(testBoart, "SNO"))
	fmt.Println(exist3(test2, "AAB"))
	fmt.Println(exist3(test3, "AAA")) // false
}

/********************使用字典树，DFS，回溯****************************/

// 用数组实现的字典树，比较浪费空间，查询效率较高
type Trie struct {
	child [58]*Trie
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
		j = word[i] - 'A'
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
		j = word[i] - 'A'
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
		j = prefix[i] - 'A'
		if cur.child[j] == nil {
			return false
		}
		cur = cur.child[j]
	}
	return true
}

// 1. 将二维数组构成一颗字典树
func exist(board [][]byte, word string) bool {
	// 将world构造字典树
	trie := Constructor()
	trie.Insert(word)
	rows := len(board)
	cols := len(board[0])
	visisted := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		visisted[i] = make([]bool, cols)
		for j := 0; j < cols; j++ {
			// 针对board[i][j]
			visisted[i][j] = false
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 针对board[i][j]
			if board[i][j] == word[0] {
				// i, j开始找
				if helper(board, i, j, visisted, "", &trie) {
					return true
				}
			}
		}
	}
	return false
}

// 从i, j开始，能插入的字典树
func helper(board [][]byte, i, j int, visited [][]bool, str string, trie *Trie) bool {
	// terminator
	// 判断越界
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 {
		return false
	}
	// 判断是否访问过
	if visited[i][j] {
		return false
	}
	// process
	str += string(board[i][j])
	if !trie.StartsWith(str) {
		return false
	}
	if trie.Search(str) {
		return true
	}

	visited[i][j] = true
	// drill down
	if helper(board, i, j+1, visited, str, trie) ||
		helper(board, i, j-1, visited, str, trie) ||
		helper(board, i+1, j, visited, str, trie) ||
		helper(board, i-1, j, visited, str, trie) {
		return true
	}
	// clear status
	visited[i][j] = false
	return false
}

/********************DFS，回溯****************************/
func exist2(board [][]byte, word string) bool {
	//dir := [4][2]int{[2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0}}
	rows := len(board)
	cols := len(board[0])
	visisted := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		visisted[i] = make([]bool, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 针对board[i][j]
			if board[i][j] == word[0] {
				// i, j开始找
				if dfs2(board, i, j, visisted, "", word) {
					return true
				}
			}
		}
	}
	return false
}

func dfs2(board [][]byte, i, j int, visited [][]bool, str string, target string) bool {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[0])-1 {
		return false
	}
	// 判断是否访问过
	if visited[i][j] {
		return false
	}
	// process
	str += string(board[i][j])
	// 判断前面部分是否匹配
	for i := 0; i < len(str); i++ {
		if str[i] != target[i] {
			return false
		}
	}
	if str == target {
		return true
	}

	visited[i][j] = true
	// drill down
	if dfs2(board, i, j+1, visited, str, target) ||
		dfs2(board, i, j-1, visited, str, target) ||
		dfs2(board, i+1, j, visited, str, target) ||
		dfs2(board, i-1, j, visited, str, target) {
		return true
	}
	// clear status
	visited[i][j] = false
	return false
}

/********************DFS，回溯，使用方向数组，代码规整****************************/
func exist3(board [][]byte, word string) bool {
	//dir := [4][2]int{[2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0}}
	rows := len(board)
	cols := len(board[0])
	visisted := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		visisted[i] = make([]bool, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 针对board[i][j]
			if board[i][j] == word[0] {
				// i, j开始找
				visisted[i][j] = true // 注意：这也得加上标记
				if dfs3(board, i, j, visisted, "", word) {
					return true
				}
				visisted[i][j] = false
			}
		}
	}
	return false
}

func dfs3(board [][]byte, i, j int, visited [][]bool, str string, target string) bool {
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}

	// terminator
	// 判断前面部分是否匹配
	if len(str) < len(target) && board[i][j] != target[len(str)] {
		return false
	}
	str += string(board[i][j])
	if str == target {
		return true
	}

	// drill down
	for k := 0; k < 4; k++ {
		newx := i + dx[k]
		newy := j + dy[k]
		// 剪枝
		if newx >= 0 && newx < len(board) && newy >= 0 && newy < len(board[0]) && !visited[newx][newy] {
			visited[newx][newy] = true
			// drill down
			if dfs3(board, newx, newy, visited, str, target) {
				return true
			}
			// clear status
			visited[newx][newy] = false
		}
	}
	return false
}
