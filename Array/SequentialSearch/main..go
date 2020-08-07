package main

import "fmt"

func main() {
	var slice []int = []int{17,33,54,67,88}
	for _,v := range slice{
		fmt.Println("******")
		fmt.Println(v)
		if v >= 54 {
			fmt.Printf("find number %d",v)
			break
		}
	}
}
