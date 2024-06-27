package LeetCode210_CourseScheduleII

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 0 {
		return []int{0}
	}

	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	result := make([]int, 0)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	qeue := []int{}

	for i, v := range indeg {
		if v == 0 {
			qeue = append(qeue, i)
		}
	}
	for len(qeue) > 0 {
		u := qeue[0]
		qeue = qeue[1:]
		result = append(result, u)
		for _, value := range edges[u] {
			indeg[value]--
			if indeg[value] == 0 {
				qeue = append(qeue, value)
			}
		}
	}
	if numCourses == len(result) {
		return result
	}
	return nil
}
