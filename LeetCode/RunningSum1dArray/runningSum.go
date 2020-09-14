package main

import "fmt"

func main() {
	var slice []int = []int{17,33,54,67,88}
	t := runningSum(slice)
	fmt.Println(t)
	n := runningSumAdvance(slice)
	fmt.Println(n)

}


func runningSum(nums []int) []int {
	slice:= make([]int,len(nums))
	sum := 0
	for i,j := range nums {
		sum = sum + j
		slice[i] = sum
	}
	return slice
}

func runningSumAdvance(nums []int) []int {
	for i:=range nums{
		if i!=0{
			nums[i]=nums[i-1]+nums[i]
		}
	}
	return nums
}
