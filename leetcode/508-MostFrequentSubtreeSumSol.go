package leetcode

import (
	"github.com/tonyliang/coding-interview-note/helper"
	"github.com/tonyliang/coding-interview-note/tree"
)

func solution508(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	counter := make(map[int]int)
	var helper508 func(*tree.TreeNode) int
	helper508 = func(node *tree.TreeNode) int {
		if node == nil {
			return 0
		}
		left := helper508(node.Left)
		right := helper508(node.Right)
		cur := left + right + node.Val
		counter[cur]++
		return cur
	}
	helper508(root)

	maxVal := helper.MinInt64
	res := []int{}
	for k, v := range counter {
		if v > maxVal {
			maxVal = v
			res = []int{k}
		} else if v == maxVal {
			res = append(res, k)
		}
	}
	return res
}
