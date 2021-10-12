package leetcode

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
)

type Question503 struct{}
type testcase503 struct {
	params []int
	ans    []int
}

func (q *Question503) Run() {
	tests := []testcase503{
		{
			params: []int{},
			ans:    []int{},
		},
		{
			params: []int{1},
			ans:    []int{-1},
		},
		{
			params: []int{1, 2, 1},
			ans:    []int{2, -1, 2},
		},
	}
	for i, t := range tests {
		res := solution503(t.params)
		if !helper.Equal(res, t.ans) {
			log.Fatalf("[Question503] test %d failed. Expect %v Got %v\n", i, t.ans, res)
		}
		fmt.Println("[Question503] All tests passed.")
	}
}
