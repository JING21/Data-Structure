package main

import "fmt"

func main() {
	var slice []int = []int{3,1,3,1,3,3,2,4}
	r := numIdenticalPairs(slice)
	fmt.Println(r)
}

func numIdenticalPairs(nums []int) int {
	var  count = 0
	var  resultmap map[int]int = make(map[int]int)
	for _, value := range nums{
		resultmap[value] += 1
	}
	for _, v := range resultmap {
		count += v*(v-1)/2
	}
	return count
}