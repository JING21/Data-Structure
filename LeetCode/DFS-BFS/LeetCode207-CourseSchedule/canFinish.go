package LeetCode207_CourseSchedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	//构建邻接表 （v是k的后修课程），key:课程1 ，value：[课程0，课程2]
	afterMap := map[int][]int{}
	//index是课程, values是课程入度
	in := make([]int, numCourses)

	for _, p := range prerequisites {
		//遍历课程情况，分别获取后修，先修课程
		after, pre := p[0], p[1]
		//构建邻接表,index是先修课程，v是后续课程的数组
		afterMap[pre] = append(afterMap[pre], after)
		in[after]++
	}

	//初始化bfs队列
	bfs := []int{}

	for courseIndex, degree := range in {
		if degree == 0 {
			bfs = append(bfs, courseIndex)
		}
	}

	//已学习过得课程数
	alreadyLearn := 0

	for len(bfs) > 0 {
		//已学习的课程数量++
		alreadyLearn++
		//bfs内是入度为0的课程
		learnedCourse := bfs[0]
		//相当于pop出为0的课程，表示学习完了
		bfs = bfs[1:]
		//遍历课程关系
		for _, afterCourse := range afterMap[learnedCourse] {
			//学完after后续课程的先修课程，所以after课程的入度--
			in[afterCourse]--
			//入度为0表示无前序课程了，可以入队
			if in[afterCourse] == 0 {
				bfs = append(bfs, afterCourse)
			}
		}
	}
	return alreadyLearn == numCourses
}

func canFinish2(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	result := make([]int, 0)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}
	//初始化队列
	q := []int{}
	for i, v := range indeg {
		if v == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		// 从头开始
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}
