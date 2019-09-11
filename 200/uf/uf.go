package uf

type UF []int

func NewUF(len int) UF {
	temp := make([]int, len)
	return UF(temp)
}

// 找到某个节点的根，即找到它的老大，并做路径压缩优化
func (uf UF) findRoot(a int) int {
	root := a
	for root != uf[root] {
		root = uf[root]
	}
	// 路径压缩优化，中间节点全部指向根节点，即全部指向老大
	for a != uf[a] {
		temp := uf[a]
		uf[a], a = root, temp
	}
	return root
}

func (uf UF) Union(a int, b int) {
	aroot := uf.findRoot(a)
	broot := uf.findRoot(b)
	uf[broot] = aroot
}

// 判断两个节点是否属于一个集合，即判断是不是有一个最终老大
func (uf UF) Connected(a int, b int) bool {
	aroot := uf.findRoot(a)
	broot := uf.findRoot(b)
	return aroot == broot
}

func (uf UF) ParentCount() int {
	// 使用Map来计数，找根
	m := make(map[int]bool)
	ok, res := false, 0
	for i := 0; i < len(uf); i++ {

		if _, ok = m[uf.findRoot(i)]; ok {
			continue
		}
		res++
		m[uf[i]] = true
	}
	return res
}
