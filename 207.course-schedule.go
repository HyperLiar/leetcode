/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	return dfs(numCourses, prerequisites)
}

func dfs(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	visit := make([]int, numCourses)

	for _, a := range prerequisites {
		graph[a[1]] = append(graph[a[1]], a[0])
	}

	for i := 0; i < numCourses; i++ {
		if !canFinishDFS(graph, &visit, i) {
			return false
		}
	}

	return true
}

func canFinishDFS(graph [][]int, visit *[]int, i int) bool {
	if (*visit)[i] == -1 {
		return false
	} else if (*visit)[i] == 1 {
		return true
	}

	(*visit)[i] = -1
	for _, a := range graph[i] {
		if !canFinishDFS(graph, visit, a) {
			return false
		}
	}

	(*visit)[i] = 1
	return true
}

func bfs(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	in := make([]int, numCourses)
	queue := []int{}

	for _, a := range prerequisites {
		graph[a[1]] = append(graph[a[1]], a[0])
		in[a[0]]++
	}

	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		for _, a := range graph[first] {
			in[a]--
			if in[a] == 0 {
				queue = append(queue, a)
			}
		}
	}

	for i := 0; i < numCourses; i++ {
		if in[i] > 0 {
			return false
		}
	}

	return true
}

// @lc code=end

