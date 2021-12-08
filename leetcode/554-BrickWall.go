package leetcode

import (
	"fmt"
	"log"
)

type Question554 struct{}
type testcase554 struct {
	wall [][]int
	ans  int
}

func (q *Question554) Run() {
	tests := []testcase554{
		// {
		// 	wall: [][]int{{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}},
		// 	ans:  2,
		// },
		// {
		// 	wall: [][]int{{1}, {1}, {1}},
		// 	ans:  3,
		// },
		{
			wall: [][]int{{1, 1}, {2}, {1, 1}},
			ans:  1,
		},
	}
	for i, t := range tests {
		res := solution554(t.wall)
		if res != t.ans {
			log.Fatalf("[Question 554] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question 554] All tests passed")
}
