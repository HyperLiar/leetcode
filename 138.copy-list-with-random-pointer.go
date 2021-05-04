/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	temp, tempNew := head, &Node{}
	newHead := tempNew
	m := make(map[*Node]*Node)

	for temp != nil {
		tempNew.Next = &Node{Val:temp.Val}
		m[temp] = tempNew.Next

		// move forward
		temp = temp.Next
		tempNew = tempNew.Next
	}
	
	temp, tempNew = head, newHead.Next
	for temp != nil {
		if temp.Random != nil {
			tempNew.Random = m[temp.Random]
		}
		temp = temp.Next
		tempNew = tempNew.Next
	}

	return newHead.Next
}
// @lc code=end

