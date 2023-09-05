package LeetCode93_RestoreIPAddresses_

import (
	"strconv"
	"strings"
)

var (
	path   []string
	result []string
)

func restoreIpAddresses(s string) []string {
	path, result = make([]string, 0, len(s)), make([]string, 0)
	backTracking(s, 0)
	return result
}

func backTracking(s string, start int) {
	if len(path) == 4 {
		if start == len(s) {
			str := strings.Join(path, ".")
			result = append(result, str)
		}
		return
	}
	for i := start; i < len(s); i++ {
		if i != start && s[start] == '0' {
			break
		}
		str := s[start : i+1]
		num, _ := strconv.Atoi(str)
		if num >= 0 && num <= 255 {
			path = append(path, str)
			backTracking(s, i+1)
			path = path[:len(path)-1]
		} else {
			break
		}
	}
}
