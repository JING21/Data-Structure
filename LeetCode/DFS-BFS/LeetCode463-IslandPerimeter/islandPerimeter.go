package LeetCode463_IslandPerimeter

func islandPerimeter(grid [][]int) int {
	var land, border int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				land++

				if i < len(grid)-1 && grid[i+1][j] == 1 {
					border++
				}
				if j < len(grid[0])-1 && grid[i][j+1] == 1 {
					border++
				}

			}

		}
	}

	return 4*land - 2*border
}

func islandPerimeter2(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return dfs(grid, i, j)
			}
		}
	}
	return 0
}

func dfs(grid [][]int, r, c int) int {
	if !inArea(grid, r, c) {
		return 1
	}

	if grid[r][c] == 0 {
		return 1
	}

	if grid[r][c] != 1 {
		return 0
	}

	grid[r][c] = 2

	return dfs(grid, r-1, c) + dfs(grid, r+1, c) + dfs(grid, r, c-1) + dfs(grid, r, c+1)
}

func inArea(grid [][]int, r, c int) bool {
	return 0 <= r && r < len(grid) && c >= 0 && c < len(grid[0])
}
