package LeetCode348_DesignTicTacToe

type TicTacToe struct {
	//矩阵大小
	number int
	//每行之和
	rowSum []int
	//每列之和
	columnSum []int
	//对角线之和
	diagonalSum, antiDiagonalSum int
}

func Constructor(n int) TicTacToe {
	newTic := TicTacToe{
		number:          n,
		rowSum:          make([]int, n),
		columnSum:       make([]int, n),
		diagonalSum:     0,
		antiDiagonalSum: 0,
	}
	return newTic
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	if player == 1 {
		this.rowSum[row]++
		this.columnSum[col]++
		if row == col {
			this.diagonalSum++
		}
		if row+col == this.number-1 {
			this.antiDiagonalSum++
		}
	}
	if player == 2 {
		this.rowSum[row]--
		this.columnSum[col]--
		if row == col {
			this.diagonalSum--
		}
		if row+col == this.number-1 {
			this.antiDiagonalSum--
		}
	}

	//如果行，列，对角线，反对角线的和为矩阵大小则表示连成一片
	if this.rowSum[row] == this.number || this.columnSum[col] == this.number || this.diagonalSum == this.number || this.antiDiagonalSum == this.number {
		return 1
	}

	if this.rowSum[row] == -this.number || this.columnSum[col] == -this.number || this.diagonalSum == -this.number || this.antiDiagonalSum == -this.number {
		return 2
	}

	return 0
}
