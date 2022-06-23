package offer40_getKmin

func getLeastNumbers3(arr []int, k int) []int {
	if k==0{//k=0的情况不能忘记
		return []int{}
	}
	start, end := 0, len(arr)-1
	index := partition(arr, start, end)
	for index != k-1 {//一旦index==k-1，即可停止。
		if index > k-1 {
			end = index - 1
		} else {
			start = index + 1
		}
		index = partition(arr, start, end)
	}
	return arr[:k]//注意：返回的这k个元素不一定是有序的
}

func partition(arr []int, start, end int) int {
	//pir := arr[l]
	//for l < h {
	//	for l < h && arr[h] >= pir {
	//		h--
	//	}
	//	arr[l] = arr[h]
	//	for l < h && arr[l] <= pir {
	//		l++
	//	}
	//	arr[h] = arr[l]
	//}
	//arr[l] = pir
	//return l

	pivot := arr[end]

	var i = start
	for j := start; j < end; j++{
		if arr[j] < pivot{
			if !(i == j){
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]

	return  i
}


