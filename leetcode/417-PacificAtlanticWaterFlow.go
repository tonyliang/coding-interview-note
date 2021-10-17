package leetcode

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
)

type Question417 struct{}

type testcase417 struct {
	heights [][]int
	ans     [][]int
}

func (q *Question417) Run() {
	tests := []testcase417{
		{
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			ans: [][]int{
				{0, 4},
				{1, 3},
				{1, 4},
				{2, 2},
				{3, 0},
				{3, 1},
				{4, 0},
			},
		},
		{
			heights: [][]int{
				{2, 1},
				{1, 2},
			},
			ans: [][]int{
				{0, 0},
				{0, 1},
				{1, 0},
				{1, 1},
			},
		},
	}
	for i, t := range tests {
		res := solution417(t.heights)
		if !helper.Equal(res, t.ans) {
			log.Fatalf("[Question417] test %d failed, expect: %v got %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question417] All tests passed")
}
