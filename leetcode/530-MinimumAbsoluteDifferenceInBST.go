package leetcode

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/tree"
)

type Question530 struct{}
type testcase530 struct {
	root *tree.TreeNode
	ans  int
}

func (q *Question530) Run() {
	tests := []testcase530{
		{
			root: BuildTree([]int{4, 2, 6, 1, 3}),
			ans:  1,
		},
		{
			root: BuildTree([]int{1, 0, 48, TreeNULL, TreeNULL, 12, 49}),
			ans:  1,
		},
		{
			root: BuildTree([]int{90, 69, TreeNULL, 49, 89, TreeNULL, 52}),
			ans:  1,
		},
	}
	for i, t := range tests {
		res := solution530(t.root)
		if res != t.ans {
			log.Fatalf("[Question530] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question530] All tests passed")
}
