package main

import "fmt"



func BinarySearch(slice []int, num int)(result bool, mid int, count int ){
	start := 0
	end := len(slice) - 1
	mid = (start+end) / 2
	for count = 1; count <= len(slice); count++{
		if slice[mid] == num {
			result = true
			return
		}else if num > slice[mid]{
			start = mid + 1
		}else {
			end = mid - 1
		}
		mid = (start+end) / 2
		fmt.Printf("start:%v, end:%v, middle:%v\n", start, end, mid)
	}
	return
}


func main() {
	l := make([]int, 100)
	for i := 1; i <= 100; i++ {
		l[i-1] = i
	}
	key := 100
	result, index, count := BinarySearch(l,key)
	fmt.Printf("search key:%v, result:%v, index:%v, count:%v\n", key, result, index, count)
}
