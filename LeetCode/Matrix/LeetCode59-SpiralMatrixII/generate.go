package LeetCode59_SpiralMatrixII

func generateMatrix(n int) [][]int {
	top := 0
	bottom := n - 1
	left := 0
	right := n - 1
	nums := 1
	largest := n * n
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for nums <= largest {
		for i := left; i <= right; i++ {
			matrix[top][i] = nums
			nums++
		}
		top++

		for i := top; i <= bottom; i++ {
			matrix[i][right] = nums
			nums++
		}
		right--

		for i := right; i >= left; i-- {
			matrix[bottom][i] = nums
			nums++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			matrix[i][left] = nums
			nums++
		}
		left++
	}

	return matrix
}
