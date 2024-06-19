package LeetCode695_MaxAreaOfIsland

func maxAreaOfIsland(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				m := dfs(grid, i, j)
				result = max(result, m)
			}
		}
	}
	return result
}

func dfs(grid [][]int, r, c int) int {

	if !inArea(grid, r, c) {
		return 0
	}

	if grid[r][c] != 1 {
		return 0
	}

	grid[r][c] = 2

	return 1 + dfs(grid, r-1, c) + dfs(grid, r+1, c) + dfs(grid, r, c-1) + dfs(grid, r, c+1)
}

func inArea(grid [][]int, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
