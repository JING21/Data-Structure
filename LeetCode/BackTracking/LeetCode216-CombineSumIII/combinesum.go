package LeetCode216_CombineSumIII

var (
	result [][]int
	path   []int
)

func combinationSum3(k int, n int) [][]int {
	result, path = make([][]int, 0), make([]int, 0, k)
	backTracking(k, n, 1, 0)
	return result
}

func backTracking(k int, n int, start int, sum int) {
	if len(path) == k {
		if sum == n {
			tmp := make([]int, k)
			copy(tmp, path)
			result = append(result, tmp)
		}
		return
	}

	for i := start; i <= 9; i++ {
		if sum+i > n || 9-i+1 < k-len(path) {
			break
		}
		path = append(path, i)
		backTracking(k, n, i+1, sum+i)
		path = path[:len(path)-1]
	}
}
