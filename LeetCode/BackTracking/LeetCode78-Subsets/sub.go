package LeetCode78_Subsets

var (
	path   []int
	result [][]int
)

func subsets(nums []int) [][]int {
	result, path = make([][]int, 0), make([]int, 0, len(nums))
	backTracking(nums, 0)
	return result
}

func backTracking(nums []int, start int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	result = append(result, tmp)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		backTracking(nums, i+1)
		path = path[:len(path)-1]
	}
}
