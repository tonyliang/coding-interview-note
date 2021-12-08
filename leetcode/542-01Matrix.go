package leetcode

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
)

type Question542 struct{}
type testcase542 struct {
	matrix [][]int
	ans    [][]int
}

func (q *Question542) Run() {
	tests := []testcase542{
		{
			matrix: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			ans:    [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		},
		{
			matrix: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}},
			ans:    [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		},
		{
			matrix: [][]int{{1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 0, 0, 1, 1}, {1, 1, 0, 0, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}},
			ans:    [][]int{{4, 3, 2, 2, 3, 4}, {3, 2, 1, 1, 2, 3}, {2, 1, 0, 0, 1, 2}, {2, 1, 0, 0, 1, 2}, {3, 2, 1, 1, 2, 3}, {4, 3, 2, 2, 3, 4}},
		},
	}
	for i, t := range tests {
		res := solution542(t.matrix)
		if !helper.Equal(res, t.ans) {
			log.Fatalf("[Question 542] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question 542] All tests passed")
}
