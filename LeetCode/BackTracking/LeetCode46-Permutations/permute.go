package LeetCode46_Permutations

var (
	result [][]int
	path   []int
	used   []bool
)

func permute(nums []int) [][]int {
	result, path = make([][]int, 0), make([]int, 0, len(nums))
	used = make([]bool, len(nums))
	backTracking()
	return result
}
