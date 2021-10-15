package QuickSort

func quickSort(a []int){
	n := len(a)
	
	if n < 2{
		return
	}
	
	for i := 1; i < n; i++{
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			swap(a, j, j-1)
		}
	}
}


func swap(slice []int, i,j int){
	slice[i], slice[j] = slice[j], slice[i]
}