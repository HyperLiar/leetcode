/*
 * @lc app=leetcode id=752 lang=golang
 *
 * [752] Open the Lock
 */

// @lc code=start
func openLock(deadends []string, target string) int {
	start := "0000"
	dead := map[string]bool{}
	for _, deadend := range deadends {
		dead[deadend] = true
	}
	if _, ok := dead[start]; ok {
		return -1
	}
	if start == target {
		return 0
	}

	q1, q2 := []string{start}, []string{target}
	m1, m2 := map[string]bool{start: true}, map[string]bool{target: true}
	s1, s2 := 0, 0

	for len(q1) != 0 && len(q2) != 0 {
		s1++
		size := len(q1)
		for k := 0; k < size; k++ {
			first := q1[k]
			for i := 0; i < 4; i++ {
				for j := -1; j <= 1; j += 2 {
					next := first
					next = turn(next, i, j)
					if m2[next] {
						return s1 + s2
					}
					if m1[next] || dead[next] {
						continue
					}
					q1 = append(q1, next)
					m1[next] = true
				}
			}
		}
		q1 = q1[size:]
		q1, q2 = q2, q1
		m1, m2 = m2, m1
		s1, s2 = s2, s1
	}

	return -1
}

func turn(s string, i, j int) string {
	b := []byte(s)
	v := byte((int(b[i]-'0')+10+j)%10 + '0')
	b[i] = v

	return string(b)
}

// @lc code=end

