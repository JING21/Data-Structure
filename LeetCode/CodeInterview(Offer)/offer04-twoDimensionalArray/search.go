package offer04_twoDimensionalArray


func findNumberIn2DArray(matrix [][]int, target int) bool {
	i := len(matrix)-1
	j := 0
	for i >= 0 {
		if j < len(matrix[i]) {
			if target < matrix[i][j] {
				i--
			} else if target > matrix[i][j] {
				j++
			} else if target == matrix[i][j] {
				return true
			}
		}else {
				return false
			}
		}
	return false
}


