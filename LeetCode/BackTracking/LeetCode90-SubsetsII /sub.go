package LeetCode90_SubsetsII_

import "sort"

var (
	path   []int
	result [][]int
)

func subsetWithDup(nums []int) [][]int {
	path, result = make([]int, 0, len(nums)), make([][]int, 0)
	sort.Ints(nums)
	backTracking(nums, 0)
	return result
}

func backTracking(nums []int, start int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	result = append(result, tmp)

	for i := start; i < len(nums); i++ {
		if i != start && nums[i] == nums[i-1] {
			continue
		}

		path = append(path, nums[i])
		backTracking(nums, i+1)
		path = path[:len(path)-1]
	}
}
