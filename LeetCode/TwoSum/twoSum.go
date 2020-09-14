package main

import "fmt"

func main() {
	var slice []int = []int{3,2,4}
	r := twoSum(slice,6)
	fmt.Println(r)
}



func twoSum(nums []int, target int) []int {
	result := []int{}
	m := make(map[int]int)
	for i,k := range nums {
		fmt.Println("i is",i)
		fmt.Println("k is",k)
		fmt.Println("**********")
		fmt.Println("map is",m)
		if value,exist := m[target-k];exist {
			fmt.Println("**********")
			fmt.Println("value is",value)
			fmt.Println("**********")
			fmt.Println("exist is",exist)
			result = append(result,value)
			result = append(result,i)
		}
		m[k] = i
		//fmt.Println(m)
	}
	return result
}