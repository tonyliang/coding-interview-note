package leetcode

import (
	"fmt"
	"log"
)

type Question507 struct{}
type testcase507 struct {
	num int
	ans bool
}

func (q *Question507) Run() {
	tests := []testcase507{
		{
			num: 28,
			ans: true,
		},
		{
			num: 6,
			ans: true,
		},
		{
			num: 496,
			ans: true,
		},
		{
			num: 8128,
			ans: true,
		},
		{
			num: 2,
			ans: false,
		},
		{
			num: 99999995,
			ans: false,
		},
	}
	for i, t := range tests {
		res := solution507(t.num)
		if res != t.ans {
			log.Fatalf("[Question507] test %d failed, expect %v got %v", i, t.ans, res)
		}
		fmt.Println("[Question507] All tests passed")
	}
}
