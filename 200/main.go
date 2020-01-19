package main

import (
	"fmt"
)

func main() {
	input1 := [][]byte{[]byte{'1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '0', '1', '1'}, []byte{'0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '0'}, []byte{'1', '0', '1', '1', '1', '0', '0', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, []byte{'1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, []byte{'1', '0', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, []byte{'1', '0', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1'}, []byte{'0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '0', '1', '1', '1', '1'}, []byte{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '0', '1', '1'}, []byte{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, []byte{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, []byte{'0', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, []byte{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, []byte{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, []byte{'1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1'}, []byte{'1', '0', '1', '1', '1', '1', '1', '0', '1', '1', '1', '0', '1', '1', '1', '1', '0', '1', '1', '1'}, []byte{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '1', '1', '0'}, []byte{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', '1', '1', '1', '0', '0'}, []byte{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, []byte{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}}
	fmt.Println("input1", len(input1), len(input1[0]))
	input2 := [][]byte{[]byte{'1', '1', '0', '0', '0'}, []byte{'1', '1', '0', '0', '0'}, []byte{'0', '0', '1', '0', '0'}, []byte{'0', '0', '0', '1', '1'}}
	input3 := [][]byte{[]byte{'1'}}
	input4 := [][]byte{[]byte{'0'}, []byte{'0'}, []byte{'0'}}
	input5 := [][]byte{[]byte{'1', '0', '0'}, []byte{'0', '0', '0'}, []byte{'0', '0', '1'}}
	input6 := [][]byte{[]byte{'0', '0', '0', '1'}}
	fmt.Println(numIslands(input1))
	fmt.Println(numIslands(input2))
	fmt.Println(numIslands(input3))
	fmt.Println(numIslands(input4))
	fmt.Println(numIslands(input5))
	fmt.Println(numIslands(input6))
	fmt.Println()
	fmt.Println(numIslands2(input1))
	fmt.Println(numIslands2(input2))
	fmt.Println(numIslands2(input3))
	fmt.Println(numIslands2(input4))
	fmt.Println(numIslands2(input5))
	fmt.Println(numIslands2(input6))
}

/*********并查集实现*********/

type UF struct {
	count int
	data  []int
}

func MakeUF(grid [][]byte) *UF {
	nr := len(grid)
	nc := len(grid[0])
	temp := make([]int, nr*nc)
	cnt := 0
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				temp[i*nc+j] = i*nc + j
				cnt++
			}
		}
	}
	return &UF{
		count: cnt,
		data:  temp,
	}
}

func (uf *UF) Union(p int, q int) {
	rootA := uf.FindRoot(p)
	rootB := uf.FindRoot(q)
	if rootA != rootB {
		uf.data[rootA] = rootB //union
		uf.count--
	}
}

// 找到根并做路径压缩
func (uf *UF) FindRoot(p int) int {
	root := p
	// 先找到根
	for root != uf.data[root] {
		root = uf.data[root]
	}
	// 找到根后做路径压缩，将路径上的点都指向根
	temp := p
	for temp != uf.data[temp] {
		par := uf.data[temp]
		uf.data[temp], temp = root, par
	}
	return root
}

func (uf *UF) Connected(p int, q int) bool {
	return uf.FindRoot(p) == uf.FindRoot(q)
}

func (uf *UF) Count() int {
	return uf.count
}

/******************/

// 并查集
func numIslands(grid [][]byte) int {
	// 将网格转换成一维
	nr := len(grid)
	if nr == 0 {
		return 0
	}
	nc := len(grid[0])
	uf := MakeUF(grid) // 用grid初始化

	// 四联通方向数组
	row, col := []int{-1, 1, 0, 0}, []int{0, 0, 1, -1}
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				for k := 0; k < 4; k++ {
					if i+row[k] >= 0 && i+row[k] < nr && j+col[k] >= 0 && j+col[k] < nc && grid[(i + row[k])][(j+col[k])] == '1' {
						// 合并
						uf.Union(i*nc+j, (i+row[k])*nc+(j+col[k]))
					}
				}
			}
		}
	}
	return uf.Count()
}

// BFS
func numIslands2(grid [][]byte) int {
	nr := len(grid)
	if nr == 0 {
		return 0
	}
	nc := len(grid[0])
	res := 0
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				res = res + 1
				queue := make([]int, 0)
				// 传入索引
				queue = append(queue, i*nc+j)
				for len(queue) != 0 {
					fmt.Println("qulen:", len(queue))
					temp := queue[0]
					queue = queue[1:]
					row := temp / nc
					col := temp % nc
					// 将相邻的都置成0
					if row-1 >= 0 && grid[row-1][col] == '1' {
						queue = append(queue, (row-1)*nc+col)
						grid[row-1][col] = '0'
					}
					if row+1 < nr && grid[row+1][col] == '1' {
						queue = append(queue, (row+1)*nc+col)
						grid[row+1][col] = '0'
					}
					if col-1 >= 0 && grid[row][col-1] == '1' {
						queue = append(queue, row*nc+col-1)
						grid[row][col-1] = '0'
					}
					if col+1 < nc && grid[row][col+1] == '1' {
						queue = append(queue, row*nc+col+1)
						grid[row][col+1] = '0'
					}
				}
			}
		}
	}
	return res
}
