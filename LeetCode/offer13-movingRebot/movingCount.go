package offer13_movingRebot

import "container/list"

var m1, n1, k1 int
var visited [][]bool

//DFS-BFS
func movingCount(m int, n int, k int) int {
	m1, n1, k1 = m, n, k

	visited = make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}

	return dfs(0, 0, 0, 0)

}

func dfs(i, j, sumi, sumj int) int {

	if i >= m1 || j >= n1 || sumi+sumj > k1 || visited[i][j] {
		return 0
	}
	visited[i][j] = true

	sumi1 := sumi + 1
	sumj1 := sumj + 1

	if (j+1)%10 == 0 {
		sumj1 = sumj - 8
	}

	if (i+1)%10 == 0 {
		sumi1 = sumi - 8
	}

	return 1 + dfs(i, j+1, sumi, sumj1) + dfs(i+1, j, sumi1, sumj)

}

//BFS
func movingCount2(m int, n int, k int) int {
	queue := list.List{}

	visited := make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}

	queue.PushBack([]int{0, 0, 0, 0})
	result := 0

	for queue.Len() > 0 {
		node := queue.Back()
		queue.Remove(node)

		bfs := node.Value.([]int)

		i := bfs[0]
		j := bfs[1]
		sumi := bfs[2]
		sumj := bfs[3]

		if i >= m || j >= n || sumi+sumj > k || visited[i][j] {
			continue
		}

		result++
		visited[i][j] = true

		sumi1 := sumi + 1
		sumj1 := sumj + 1

		if (i+1)%10 == 0 {
			sumi1 = sumi - 8
		}

		if (j+1)%10 == 0 {
			sumj1 = sumj - 8
		}

		queue.PushBack([]int{i + 1, j, sumi1, sumj})
		queue.PushBack([]int{i, j + 1, sumi, sumj1})
	}
	return result
}
