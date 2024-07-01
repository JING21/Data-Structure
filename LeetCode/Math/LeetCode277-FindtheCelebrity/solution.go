package LeetCode277_FindtheCelebrity

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		result := 0

		for i := 1; i < n; i++ {
			//当knows(i,j)为true时，表示i不是名人，当knows(i，j)为false时，表示j不是名人，所以两两对比，得到一个候选人
			if knows(result, i) {
				result = i
			}
		}
		for i := 0; i < n; i++ {
			//排除自己
			if result == i {
				continue
			}
			//如果候选人认识某个人或者某个人不认识候选人，那么候选人就不是名人,反之候选人就是名人
			if knows(result, i) || !knows(i, result) {
				return -1
			}
		}
		return result
	}
}
