package LeetCode200_NumberOfIslands

func numIslands(grid [][]byte) int {
	result := 0
	// i 是行
	for i := 0; i < len(grid); i++ {
		// j 是列
		for j := 0; j < len(grid[i]); j++ {
			// 取出陆地
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				result++
			}
		}
	}
	return result
}

func dfs(grid [][]byte, r, c int) {

	if !inArea(grid, r, c) {
		return
	}

	if grid[r][c] != '1' {
		return
	}

	grid[r][c] = '2'

	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

func inArea(grid [][]byte, r, c int) bool {
	return r < len(grid) && r >= 0 && c < len(grid[0]) && c >= 0
}
