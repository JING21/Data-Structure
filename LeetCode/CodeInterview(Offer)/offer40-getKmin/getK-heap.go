package offer40_getKmin


import (
	"container/heap"
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
// 为了实现大根堆，Less在大于时返回小于
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
// 大根堆
func getLeastNumbers2(arr []int, k int) []int {
	if k==0{return []int{}}
	h := make(IntHeap,k)
	hp:=&h
	copy(h ,IntHeap(arr[:k+1]))
	heap.Init(hp)
	for i:=k;i<len(arr);i++{
		if arr[i]<h[0]{
			heap.Pop(hp)
			heap.Push(hp,arr[i])
		}
	}
	return h

}

func main() {
	arr := []int{0,1,2,1}
	getLeastNumbers2(arr, 1)
}