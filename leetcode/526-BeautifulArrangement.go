package leetcode

import (
	"fmt"
	"log"
)

type Question526 struct{}
type testcase526 struct {
	N   int
	ans int
}

func (q *Question526) Run() {
	tests := []testcase526{

		{
			N:   1,
			ans: 1,
		},

		{
			N:   2,
			ans: 2,
		},

		{
			N:   3,
			ans: 3,
		},

		{
			N:   4,
			ans: 8,
		},

		{
			N:   5,
			ans: 10,
		},

		{
			N:   6,
			ans: 36,
		},

		{
			N:   7,
			ans: 41,
		},

		{
			N:   8,
			ans: 132,
		},

		{
			N:   9,
			ans: 250,
		},

		{
			N:   10,
			ans: 700,
		},

		{
			N:   11,
			ans: 750,
		},

		{
			N:   12,
			ans: 4010,
		},

		{
			N:   13,
			ans: 4237,
		},

		{
			N:   14,
			ans: 10680,
		},

		{
			N:   15,
			ans: 24679,
		},

		{
			N:   16,
			ans: 87328,
		},

		{
			N:   17,
			ans: 90478,
		},

		{
			N:   18,
			ans: 435812,
		},
	}
	for i, t := range tests {
		res := solution526(t.N)
		if res != t.ans {
			log.Fatalf("[Question526] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question526] All tests passed")
}
