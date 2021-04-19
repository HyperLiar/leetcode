/*
 * @lc app=leetcode id=721 lang=golang
 *
 * [721] Accounts Merge
 */

// @lc code=start
func accountsMerge(accounts [][]string) [][]string {
	emailToAccountId := make(map[string]int)

	// union
	uf := newUnionFind(len(accounts))
	for accountId, emails := range accounts {
		for i := 1; i < len(emails); i++ {
			if id, ok := emailToAccountId[emails[i]]; ok {
				uf.union(id, accountId)
			} else {
				emailToAccountId[emails[i]] = accountId
			}
		}
	}

	// find
	for i := 0; i < len(accounts); i++ {
		uf.parent[i] = uf.find(i)
	}

	emailListMap := make(map[int][]string)
	for email, accountId := range emailToAccountId {
		rootAccount := uf.find(accountId)
		if _, ok := emailListMap[rootAccount]; !ok {
			emailListMap[rootAccount] = []string{accounts[rootAccount][0]}
		}

		emailListMap[rootAccount] = append(emailListMap[rootAccount], email)
	}
	// make result
	res := make([][]string, len(emailListMap))
	resIdx := 0
	for _, emailList := range emailListMap {
		sort.Strings(emailList[1:])
		res[resIdx] = emailList
		resIdx++
	}

	return res
}

type UnionFind struct {
	size   int
	parent []int
	rank   []int
}

func newUnionFind(n int) UnionFind {
	uf := UnionFind{size: n}
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)

	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}

	return uf
}

// union p q
func (uf *UnionFind) union(p, q int) {
	rootP, rootQ := uf.find(p), uf.find(q)

	if rootP == rootQ {
		return
	}

	if uf.rank[rootP] > uf.rank[rootQ] {
		uf.parent[rootQ] = rootP
		uf.rank[rootP] += uf.rank[rootQ]
	} else {
		uf.parent[rootP] = rootQ
		uf.rank[rootQ] += uf.rank[rootP]
	}
	uf.size--
}

// find root of p
func (uf *UnionFind) find(p int) int {
	for uf.parent[p] != p {
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}

	return p
}

// @lc code=end

