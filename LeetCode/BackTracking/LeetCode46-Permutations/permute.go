package LeetCode46_Permutations

var (
	result [][]int
	path   []int
	used   []bool
)

func permute(nums []int) [][]int {
	result, path = make([][]int, 0), make([]int, 0, len(nums))
	used = make([]bool, len(nums))
	backTracking(nums, 0)
	return result
}

func backTracking(nums []int, level int) {
	if level == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backTracking(nums, level+1)
		path = path[:len(path)-1]
		used[i] = false
	}
}
