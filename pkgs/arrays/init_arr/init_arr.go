package initarr

import (
	"math/rand"
	"sort"
	"time"
)

func InitRand(size int) []int {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	var arr = make([]int, size)

	for i := 0; i < len(arr); i++ {
		arr[i] = r.Intn(size)
	}
	return arr
}

func InitOrd(size int) []int {
	arr := InitRand(size)
	sort.Ints(arr)
	return arr
}
