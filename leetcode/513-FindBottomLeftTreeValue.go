package leetcode

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/tree"
)

type Question513 struct{}
type testcase513 struct {
	root *tree.TreeNode
	ans  int
}

func (q *Question513) Run() {
	tests := []testcase513{
		{
			root: BuildTree([]int{2, 1, 3}),
			ans:  1,
		},
		{
			root: BuildTree([]int{1, 2, 3, 4, TreeNULL, 5, 6, TreeNULL, TreeNULL, 7}),
			ans:  7,
		},
	}
	for i, t := range tests {
		res := solution513(t.root)
		if res != t.ans {
			log.Fatalf("[Question513] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question513] All tests passed")
}
