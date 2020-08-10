package main

import "fmt"

func BubbleSort(slice []int){
	fmt.Println("Before sort ",slice)
	temp := 0 //临时变量(用于做交换)
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if (slice)[j] > (slice)[j+1] {
				//交换
				temp = (slice)[j]
				(slice)[j] = (slice)[j+1]
				(slice)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后arr=", slice)
}

func main() {
	var slice []int = []int{30,55,27,19,80,77,45,21,18,44}
	BubbleSort(slice)
}
