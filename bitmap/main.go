package main

import (
	"fmt"
	"github.com/boljen/go-bitmap"
	"github.com/golang/go/src/os"
)

// 依赖github.com/boljen/go-bitmap

func main() {
	bm := bitmap.New(16)
	// 如果false，索引位置i置为0，如果为true，索引位置i置为1
	bm.Set(0, true)
	fmt.Println(bm.Get(0))
	fmt.Println(bm.Len()) // 做了冗余

	cbm := bitmap.NewConcurrent(16)
	cbm.Set(0, true)
	cbm.Get(0)
	fmt.Println(cbm.Len())
	os.OpenFile()

	sl := make([]byte, 1, 3)
	fmt.Println(len(sl), cap(sl))
}
