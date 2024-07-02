package __1pack

func pack(weight []int , n,w int)  int{
	var state [][]bool
	state[0][0] = true

	if weight[0] <= w{
		state[0][weight[0]] = true
	}
	for i := 1; i < n; i++{
		for j := 0; j <= w; j++{
			if state[i-1][j] == true{
				state[i][j] = state[i-1][j]
			}
		}
		for j := 0; j <= w-weight[i]; j++{
			if state[i-1][j] == true{
				state[i][j+weight[i]] = true
			}
		}

		}
	for i := w; i >=0; i--{
		if state[n-1][i] == true {
			return i
		}
	}
	return 0
}