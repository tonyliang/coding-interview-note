package leetcode

import "github.com/tonyliang/coding-interview-note/tree"

func solution515(root *tree.TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	var dfs func(*tree.TreeNode, int)
	dfs = func(node *tree.TreeNode, level int) {
		if node == nil {
			return
		}
		if level == len(res) {
			res = append(res, node.Val)
		}
		if node.Val > res[level] {
			res[level] = node.Val
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	return res
}
