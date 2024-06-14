package LeetCode47_PermutationUnique

import "sort"

var (
	result [][]int
	path   []int
	used   []bool
)

func permuteUnique(nums []int) [][]int {
	result, path = make([][]int, 0), make([]int, 0, len(nums))
	used = make([]bool, len(nums))
	sort.Ints(nums)
	backTracking(nums)
	return result

}

func backTracking(nums []int) {
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backTracking(nums)
		path = path[:len(path)-1]
		used[i] = false
	}
}
