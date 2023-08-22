// package LeetCode131_PalindromePartitioning_
package main

import "fmt"

var (
	path   []string
	result [][]string
)

func partition(s string) [][]string {
	path, result = make([]string, 0), make([][]string, 0)
	backTracking(s, 0)
	return result
}

func backTracking(s string, start int) {
	fmt.Println("******************")
	fmt.Printf("Result is %v, start is %v\n", result, start)
	if start == len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		return
	}

	for i := start; i < len(s); i++ {
		str := s[start : i+1]
		if isPalindrome(str) {
			path = append(path, str)
			backTracking(s, i+1)
			path = path[:len(path)-1]
			fmt.Println("##################")
			fmt.Printf("Path is %v\n", path)
		}
	}

}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	partition("aab")
}
