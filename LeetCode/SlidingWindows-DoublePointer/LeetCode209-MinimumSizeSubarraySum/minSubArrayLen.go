package LeetCode209_MinimumSizeSubarraySum

import "math"

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	minSize := math.MaxInt64
	start, end := 0, 0
	sum := 0

	for end < n {
		sum += nums[end]
		for sum >= s {
			minSize = min(minSize, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if minSize == math.MaxInt64 {
		return 0
	}

	return minSize
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
