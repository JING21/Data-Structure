package offer40_getKmin

import (
	"container/heap"
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}



func getLeastNumbers2(arr []int, k int) []int {
	heap.Init(arr)
}

