package leetcode

import (
	"math"

	"github.com/tonyliang/coding-interview-note/tree"
)

func solution530(root *tree.TreeNode) int {
	pre := -1
	ans := math.MaxInt16

	var dfs func(*tree.TreeNode)

	dfs = func(node *tree.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 {
			ans = min530(node.Val-pre, ans)
		}
		pre = node.Val
		dfs(node.Right)
	}

	dfs(root)
	return ans
}

func min530(a, b int) int {
	if a < b {
		return a
	}
	return b
}
