package leetcode

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
	"github.com/tonyliang/coding-interview-note/tree"
)

type Question515 struct{}

type testcase515 struct {
	root *tree.TreeNode
	ans  []int
}

func (q *Question515) Run() {
	tests := []testcase515{
		{
			root: BuildTree([]int{1, 3, 2, 5, 3, TreeNULL, 9}),
			ans:  []int{1, 3, 9},
		},
		{
			root: BuildTree([]int{1, 2, 3}),
			ans:  []int{1, 3},
		},
		{
			root: BuildTree([]int{1}),
			ans:  []int{1},
		},
		{
			root: BuildTree([]int{1, TreeNULL, 2}),
			ans:  []int{1, 2},
		},
		{
			root: BuildTree([]int{}),
			ans:  []int{},
		},
		{
			root: BuildTree([]int{3, 9, 20, TreeNULL, TreeNULL, 15, 7}),
			ans:  []int{3, 20, 15},
		},
	}
	for i, t := range tests {
		res := solution515(t.root)
		if !helper.Equal(res, t.ans) {
			log.Fatalf("[Question515] test %d failed, expect %v got %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question515] All tests passed")
}
