package LeetCode216_CombineSumIII

import "sort"

// leetcode submit region begin(Prohibit modification and deletion)
var (
	path   []int
	result [][]int
)

func combinationSum(candidates []int, target int) [][]int {
	result, path = make([][]int, 0), make([]int, 0, len(candidates))
	sort.Ints(candidates)
	backTracking(candidates, 0, target)
	return result
}

func backTracking(candidates []int, start int, target int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}

		path = append(path, candidates[i])
		backTracking(candidates, i, target-candidates[i])
		path = path[:len(path)-1]
	}

}

//leetcode submit region end(Prohibit modification and deletion)
