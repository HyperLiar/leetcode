/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    // if len(preorder) == 0 {
	// 	return nil
	// }

	// root := &TreeNode{Val:preorder[0]}
	
	// rootIdx := 0
	// for i := 0; i < len(preorder); i++ {
	// 	if inorder[i] == preorder[0] {
	// 		rootIdx = i
	// 		break
	// 	}
	// }

	// root.Left = buildTree(preorder[1:rootIdx+1], inorder[0:rootIdx])
	// root.Right = buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:])
	// return root

	inorderMap := make(map[int]int)
	preIdx := 0
	for i := 0; i < len(inorder); i++ {
		inorderMap[inorder[i]] = i
	}

	return arrToTree(preorder, &preIdx, 0, len(inorder)-1, inorderMap)
}

func arrToTree(preorder []int, preIdx *int, start, end int, inorderMap map[int]int) *TreeNode {
	if start > end {
		return nil
	}

	rootVal := preorder[*preIdx]
	(*preIdx)++

	root := &TreeNode{Val:rootVal}

	root.Left = arrToTree(preorder, preIdx, start, inorderMap[rootVal]-1, inorderMap)
	root.Right = arrToTree(preorder, preIdx, inorderMap[rootVal]+1, end, inorderMap)

	return root
}
// @lc code=end

