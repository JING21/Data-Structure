package offer29_printMatrix



func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1

	var result []int

	for{
		//left to right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if top > bottom{
			break
		}

		//top to bottom
		for i := top; i <= bottom; i++{
			result = append(result, matrix[i][right])
		}
		right--
		if right < left{
			break
		}

		//right to left
		for  i := right; i >= left; i--{
			result = append(result, matrix[bottom][i])
		}
		bottom--
		if top > bottom{
			break
		}

		for i := bottom; i >= top; i--{
			result = append(result, matrix[i][left])
		}
		left++
		if left > right{
			break
		}
	}

	return result
}