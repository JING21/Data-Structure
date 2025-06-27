package LeetCode2239_FindClosestNumbertoZero

import "math"

func findClosestNumber(nums []int) (res int) {
	min := math.MaxInt32
	for _, v := range nums {
		if y := abs(v); y < min || y == min && v > res {
			min, res = y, v
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
