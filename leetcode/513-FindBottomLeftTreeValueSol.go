package leetcode

import "github.com/tonyliang/coding-interview-note/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func solution513(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	maxLevel := 1
	result := root.Val
	var helper513 func(*tree.TreeNode, int)
	helper513 = func(node *tree.TreeNode, level int) {
		if node == nil {
			return
		}
		helper513(node.Left, level+1)
		helper513(node.Right, level+1)
		if level > maxLevel {
			maxLevel = level
			result = node.Val
		}
	}
	helper513(root, 1)
	return result
}
