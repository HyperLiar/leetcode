/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	//return bfs(numCourses, prerequisites)
	return dfs(numCourses, prerequisites)
}

func hasCycles(subj int, graph [][]int, visited, res *[]int) bool {
	if (*visited)[subj] == -1 {
		return false
	} else if (*visited)[subj] == 1 {
		return true
	}

	(*visited)[subj] = -1
	for _, a := range graph[subj] {
		if !hasCycles(a, graph, visited, res) {
			return false
		}
	}

	(*visited)[subj] = 1
	*res = append(*res, subj)
	return true
}
func dfs(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	visited := make([]int, numCourses)
	res := []int{}

	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
	}

	for i := 0; i < numCourses; i++ {
		if !hasCycles(i, graph, &visited, &res) {
			return []int{}
		}
	}

	for i := 0; i < numCourses/2; i++ {
		res[i], res[numCourses-i-1] = res[numCourses-i-1], res[i]
	}

	return res
}


func bfs(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	in := make([]int, numCourses)
	queue := []int{}
	queueStart := 0

	for _, a := range prerequisites {
		graph[a[1]] = append(graph[a[1]], a[0])
		in[a[0]]++
	}

	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			queue = append(queue, i)
		}
	}

	for queueStart < len(queue) {
		first := queue[queueStart]
		queueStart++
		for _, a := range graph[first] {
			in[a]--
			if in[a] == 0 {
				queue = append(queue, a)
			}
		}
	}

	if len(queue) == numCourses {
		return queue
	}

	return []int{}
}

// @lc code=end

