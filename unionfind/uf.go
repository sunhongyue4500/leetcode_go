package unionfind

type UF struct {
	count int
	data  []int
}

// 可根据情况写初始化函数
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
