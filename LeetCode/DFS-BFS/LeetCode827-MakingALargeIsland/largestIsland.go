package LeetCode827_MakingALargeIsland

func largestIsland(grid [][]int) int {
	maxArea := 0
	index := 2
	areasMap := make(map[int]int, 0)

	//遍历所有陆地求出最大面积
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			//表示是陆地,搜索周围
			if grid[i][j] == 1 {
				area := dfs(grid, i, j, index)
				areasMap[index] = area
				index++
				maxArea = max(maxArea, area)
			}
		}
	}

	if maxArea == 0 {
		return 1
	}

	//遍历所有海洋，记录陆地上下左右的情况
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				//搜寻此海洋节点的四周
				set := findNeighbour(grid, i, j)

				//没有陆地
				if len(set) == 0 {
					continue
				}
				sum := 1
				for m, _ := range set {
					sum += areasMap[m]
				}
				maxArea = max(maxArea, sum)
			}
		}
	}
	return maxArea
}

func dfs(grid [][]int, r, c, index int) int {
	//越界，结束搜索
	if !isPathValid(grid, r, c) {
		return 0
	}

	//如果不是陆地或者已经被访问
	if grid[r][c] != 1 {
		return 0
	}

	//将岛屿标记为index
	grid[r][c] = index
	return 1 + dfs(grid, r-1, c, index) + dfs(grid, r+1, c, index) + dfs(grid, r, c-1, index) + dfs(grid, r, c+1, index)
}

func findNeighbour(grid [][]int, i, j int) map[int]bool {
	set := make(map[int]bool)

	//海水左边
	if isPathValid(grid, i-1, j) && grid[i-1][j] != 0 {
		set[grid[i-1][j]] = true
	}

	//右边
	if isPathValid(grid, i+1, j) && grid[i+1][j] != 0 {
		set[grid[i+1][j]] = true
	}

	//上面
	if isPathValid(grid, i, j-1) && grid[i][j-1] != 0 {
		set[grid[i][j-1]] = true
	}

	//下面
	if isPathValid(grid, i, j+1) && grid[i][j+1] != 0 {
		set[grid[i][j+1]] = true
	}

	return set
}

func isPathValid(grid [][]int, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
