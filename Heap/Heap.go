package Heap

type Heap struct {
	a  		[]int
	n   	int
	count   int
}

//init heap
func NewHeap(capcity int) *Heap{
	heap := &Heap{}
	heap.n = capcity
	heap.a = make([]int, capcity+1)
	heap.count = 0

	return heap
}

func (h *Heap)insert(data int){
	if h.count == h.n{
		return
	}

	h.count++
	h.a[h.count] = data

	i := h.count
	parent := i/2

	for parent > 0 && h.a[parent] < h.a[i]{
			swap(h.a, parent, i)
			i = parent
			parent = i/2
	}
}

func (h *Heap)removeMax(){
	if h.count == 0{
		return
	}

	swap(h.a, 1, h.count)
	h.count--

	heapifyUpToDown(h.a, h.count)
}


func heapifyUpToDown(a []int, count int){
	for i := 1; i <= count/2;{
		maxIndex := i

		if a[i] < a[i*2]{
			maxIndex = i*2
		}

		if i*2 + 1 <= count && a[maxIndex] < a[i*2+1]{
			maxIndex = i*2+1
		}

		if maxIndex == i{
			break
		}

		swap(a, i, maxIndex)
		i = maxIndex
	}
}

func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}