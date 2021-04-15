/*
 * @lc app=leetcode id=721 lang=golang
 *
 * [721] Accounts Merge
 */

// @lc code=start
func accountsMerge(accounts [][]string) [][]string {
    idList := make(map[int]string)
	emailToId := make(map[string]int)
	uf := newUnionFind(len(accounts))

	for i := 0; i < len(accounts); i++ {
		idList[i] = accounts[i][0]
		for j := 1; j < len(accounts[i]); j++ {
			currentEmail := accounts[i][j]
			if _, ok := emailToId[i]; ok {
				uf.unify(i, emailToId[currentEmail])
			} else {
				emailToId[currentEmail] = i
			}
		}
	}

	idToEmail := make(map[int][]string)

	for email, id := range emailToId {
		sourceId := uf.find(id)

		if _, ok := emailToId[id];!ok {
			idToEmail[id] = []string{}
		}

		idToEmail[id] = append(idToEmail[id], email)
	}

	res := make([][]string, 0)

	for id, emailList := range idToEmail {
		sort.Strings(emailList)
		temp := []string{idList[id]}
		temp = append(temp, emailList...)
		res = append(res, temp)
	}

	return res
}

type UnionFind struct {
	size int
	parent []int
	rank []int
}

func newUnionFind(size int) UnionFind {
	uf := UnionFind{size:size}
	uf.parent = make([]int, size)
	uf.rank = make([]int, rank)

	return uf
}

func (uf *UnionFind)find(id int) int {
	root := id
	// find root
	for root != uf.parent[root] {
		root = uf.parent[root]
	}

	// set all node in path to root
	for id != root {
		next = uf.parent[id]
		uf.parent[id] = root
		id = next
	}

	return root
}

func (uf *UnionFind) unify(id1, id2 int) {
	root1, root2 := uf.find(id1), uf.find(id2)

	if root1 == root2 {
		return
	}

	if uf.rank[id1] < uf.rank[id2] {
		uf.parent[root1] = root2
		uf.rank[root2] += uf.rank[root1]
	} else {
		uf.parent[root2] = root1
		uf.rank[root1] += uf.rank[root2]
	}
}
// @lc code=end

