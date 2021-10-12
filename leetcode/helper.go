package leetcode

import (
	"github.com/tonyliang/coding-interview-note/helper"
	"github.com/tonyliang/coding-interview-note/tree"
)

var TreeNULL = helper.MinInt64

func BuildTree(nums []int) *tree.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &tree.TreeNode{
		Val: nums[0],
	}
	queue := []*tree.TreeNode{root}
	i := 1
	n := len(nums)
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && nums[i] != TreeNULL {
			node.Left = &tree.TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && nums[i] != TreeNULL {
			node.Right = &tree.TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func TreeToIntArray(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*tree.TreeNode{root}
	nums := []int{}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node == nil {
				nums = append(nums, TreeNULL)
			} else {
				nums = append(nums, node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}
		queue = queue[size:]
	}
	i := len(nums)
	for i > 0 && nums[i-1] == TreeNULL {
		i--
	}
	return nums[:i]
}
