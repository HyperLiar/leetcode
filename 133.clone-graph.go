/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

 func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    m := make(map[*Node]*Node)
    visited := make(map[*Node]bool)
    helper(node, &m, &visited)

    for x, y := range m {
        y.Neighbors = make([]*Node, len(x.Neighbors))
        y.Val = x.Val
        for i, z := range x.Neighbors {
            y.Neighbors[i] = m[z]
        }
    }

    return m[node]
}

func helper(node *Node, m *map[*Node]*Node, visited *map[*Node]bool) {
    if (*visited)[node] {
        return
    }
    (*visited)[node] = true
    (*m)[node] = &Node{}
    for _, n := range node.Neighbors {
        helper(n, m, visited)
    }
}
// @lc code=end

