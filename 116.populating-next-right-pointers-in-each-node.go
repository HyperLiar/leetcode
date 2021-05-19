/*
 * @lc app=leetcode id=116 lang=golang
 *
 * [116] Populating Next Right Pointers in Each Node
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
 func connect(root *Node) *Node {
    if root == nil {
		return root
	}
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		for node := leftmost; node != nil; node = node.Next {
			node.Left.Next = node.Right

			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root

    // preConnect(root)
    // postConnect(root)
    // return root
}


func preConnect(root *Node) {
    if root == nil || root.Left == nil || root.Right == nil {
        return
    }

    root.Left.Next = root.Right
    root.Right.Next = root
    preConnect(root.Left)
    preConnect(root.Right)
}

func postConnect(root *Node) {
    if root == nil || root.Left == nil || root.Right == nil {
        return
    }

    if root.Next == nil || root.Right.Next == nil || root.Right.Next.Next == nil {
        root.Right.Next = nil
    } else {
        root.Right.Next = root.Right.Next.Next.Left
    }
    postConnect(root.Left)
    postConnect(root.Right)
}

// @lc code=end

