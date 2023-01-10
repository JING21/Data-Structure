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

func minSubArrayLen2(target int, nums []int) int {
	i := 0
	l := len(nums)  // 数组长度
	sum := 0        // 子数组之和
	result := l + 1 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}
