package main

import "fmt"

func BubbleSort(slice []int){
	fmt.Println("Before sort ",slice)
	//for i := 0; i < len(slice); i++ {
	//
	//}
}

func swap(a,b,temp int){
	temp = a
	a = b
	b = temp
}

func main() {
	var slice []int = []int{17,33,54,67,88}
	fmt.Println(len(slice))
	BubbleSort(slice)
}
