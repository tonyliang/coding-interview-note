package leetcode

import (
	"fmt"
	"log"
)

type Question494 struct{}

type testcase494 struct {
	nums   []int
	target int
	ans    int
}

func (q *Question494) Run() {
	tests := []testcase494{
		{
			nums:   []int{1, 1, 1, 1, 1},
			target: 3,
			ans:    5,
		},
		{
			nums:   []int{1},
			target: 1,
			ans:    1,
		},
		{
			nums:   []int{1, 2, 1},
			target: 0,
			ans:    2,
		},
	}
	for i, t := range tests {
		res := solution494(t.nums, t.target)
		if res != t.ans {
			log.Fatalf("[Question494] test %d failed, expect %v got %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question494] All tests passed")
}
