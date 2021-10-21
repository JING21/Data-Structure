//package InsertSort
package main

import "fmt"

func insertSort(a []int){
	n := len(a)
	
	if n < 2 {
		return
	}
	
	for i := 1; i < n; i++{
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			swap(a, j, j-1)
		}
	}
}


func insertSort2(a []int){
	n := len(a)

	if n < 2 {
		return
	}

	for i := 1; i < n; i++{
		value := a[i]
		j := i-1
		for ;j >= 0; j--{
			if a[j] > value{
				a[j+1] = a[j]
			}else {
				break
			}
		}
		a[j+1] = value
	}
}


func swap(slice []int, i,j int){
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	test := []int{2,7,1,3,5}
	fmt.Println(test)
	insertSort2(test)
	fmt.Println(test)
	fmt.Println(10&9)
}