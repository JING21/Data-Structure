package LeetCode77_Combination

var (
	result [][]int
	path   []int
)

func combine(n int, k int) [][]int {
	result, path = make([][]int, 0), make([]int, 0, k)
	backtracking(n, k, 1)
	return result
}

func backtracking(n int, k int, start int) {
	if len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		result = append(result, tmp)
		return
	}

	for i := start; i <= n-(k-len(path))+1; i++ {
		path = append(path, i)
		backtracking(n, k, i+1)
		path = path[:len(path)-1]
	}

}
