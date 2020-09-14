package main

import "fmt"

func main() {
	var slice []int = []int{2,3,5,1,3}
	r := kidsWithCandies(slice,3)
	fmt.Println(r)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	for _, j := range candies{
		maxCandies = max(maxCandies,j)
	}
	result := make([]bool,len(candies))
	for i, _ := range candies{
		result[i] = candies[i]+extraCandies >= maxCandies
	}
	return result
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}