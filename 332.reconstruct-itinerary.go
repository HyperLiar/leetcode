/*
 * @lc app=leetcode id=332 lang=golang
 *
 * [332] Reconstruct Itinerary
 */

// @lc code=start
func findItinerary(tickets [][]string) []string {
	m := map[string][]string{}
	for _, ticket := range tickets {
		m[ticket[0]] = append(m[ticket[0]], ticket[1])
	}
	for k, _ := range m {
		sort.Strings(m[k])
	}

	res := []string{}
	dfs("JFK", m, &res)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func dfs(start string, m map[string][]string, res *[]string) {
	for len(m[start]) > 0 {
		v := m[start][0]
		m[start] = m[start][1:]
		dfs(v, m, res)
	}
	*res = append(*res, start)
}
// func findItinerary(tickets [][]string) []string {
// 	m := map[string][]string{}
// 	visited := len(tickets)+1
//     for _, ticket := range tickets {
// 		m[ticket[0]] = append(m[ticket[0]], ticket[1])
// 	}
// 	for k, _ := range m {
// 		sort.Strings(m[k])
// 	}

// 	road := []string{"JFK"}
// 	res := []string{}
// 	dfs("JFK", m, &res, &road, visited)
// 	return res
// }

// func dfs(start string, m map[string][]string, res, road *[]string, visited int) {
// 	if len(*res) > 0 {
// 		return
// 	}
// 	if len(*road) == visited {
// 		*res = append([]string{}, *road...)
// 		return
// 	}
	
// 	if len(m[start]) == 0 {
// 		return
// 	}
	
// 	for i, end := range m[start] {
// 		if len(end) != 3 {
// 			continue
// 		}
// 		*road = append(*road, end)
// 		t := m[start][i]
// 		m[start][i] = ""
// 		dfs(end, m, res, road, visited)
// 		m[start][i] = t
// 		*road = (*road)[0:len(*road)-1]
// 	}
// }
// @lc code=end

