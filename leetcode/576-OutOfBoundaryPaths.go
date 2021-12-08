package leetcode

import (
	"fmt"
	"log"
)

type Question576 struct{}
type testcase576 struct {
	m           int
	n           int
	maxMove     int
	startRow    int
	startColumn int
	ans         int
}

func (q *Question576) Run() {
	tests := []testcase576{
		{
			m:           2,
			n:           2,
			maxMove:     2,
			startRow:    0,
			startColumn: 0,
			ans:         6,
		},

		{
			m:           1,
			n:           3,
			maxMove:     3,
			startRow:    0,
			startColumn: 1,
			ans:         12,
		},
	}
	for i, t := range tests {
		res := solution576(t.m, t.n, t.maxMove, t.startRow, t.startColumn)
		if res != t.ans {
			log.Fatalf("[Question 576] test %d failed, expect: %d, got: %d\n", i, t.ans, res)
		}
	}
	fmt.Println("[Question 576] All tests passed")
}
