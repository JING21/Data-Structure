package main

import "fmt"

func SelectSort(slice []int){
	for i :=0; i<len(slice)-1;i++{
		//假设slice[i]为最大值
		max := slice[i]
		maxIndex := i
		//遍历i+1之后的所有值进行比较
		for j :=i+1; j< len(slice);j++ {
			if max < slice[j] {
				max = slice[j]
				maxIndex = j
			}
		}
		// 交换
		if maxIndex != i {
			slice[i], slice[maxIndex] = slice[maxIndex], slice[i]

		}
		fmt.Printf("第%d次 %v\n", i+1, slice)
	}
}


func main() {
	arr := []int{10, 34, 19, 100, 80}
	fmt.Println(arr)
	SelectSort(arr)
	fmt.Println("main函数")
	fmt.Println(arr)

}
