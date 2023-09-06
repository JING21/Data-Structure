package LeetCode_491_NonDecreasingSubsequences

var (
	path   []int
	result [][]int
)

func findSubsequences(nums []int) [][]int {
	result, path = make([][]int, 0), make([]int, 0, len(nums))
	backTracking(nums, 0)
	return result
}

func backTracking(nums []int, start int) {
	if len(path) >= 2 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
	}

	used := make(map[int]bool, len(nums))

	for i := start; i < len(nums); i++ {
		if used[nums[i]] {
			continue
		}

		if len(path) == 0 || nums[i] >= path[len(path)-1] {
			path = append(path, nums[i])
			used[nums[i]] = true
			backTracking(nums, i+1)
			path = path[:len(path)-1]
		}

	}
}
