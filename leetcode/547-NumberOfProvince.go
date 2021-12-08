package leetcode

import (
	"fmt"
	"log"
)

type Question547 struct{}
type testcase547 struct {
	M   [][]int
	ans int
}

func (q *Question547) Run() {
	tests := []testcase547{
		{
			M:   [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			ans: 3,
		},

		{
			M:   [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
			ans: 2,
		},

		{
			M:   [][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}},
			ans: 1,
		},
	}
	for i, t := range tests {
		res := solution547(t.M)
		if res != t.ans {
			log.Fatalf("[Question 547] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question 547] All tests passed")
}
