package leetcode

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
)

type Question218 struct{}
type testcase218 struct {
	buildings [][]int
	ans       [][]int
}

func (q *Question218) Run() {
	tests := []testcase218{
		{
			buildings: [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}},
			ans:       [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}},
		},
		{
			buildings: [][]int{{1, 2, 1}, {1, 2, 2}, {1, 2, 3}, {2, 3, 1}, {2, 3, 2}, {2, 3, 3}},
			ans:       [][]int{{1, 3}, {3, 0}},
		},

		{
			buildings: [][]int{{4, 9, 10}, {4, 9, 15}, {4, 9, 12}, {10, 12, 10}, {10, 12, 8}},
			ans:       [][]int{{4, 15}, {9, 0}, {10, 10}, {12, 0}},
		},
	}
	for i, t := range tests {
		res := solution218(t.buildings)
		if !helper.Equal(res, t.ans) {
			log.Fatalf("[Question 218] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question 218] All tests passed")
}
