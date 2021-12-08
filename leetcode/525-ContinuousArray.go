package leetcode

import (
	"fmt"
	"log"
)

type Question525 struct{}
type testcase525 struct {
	nums []int
	ans  int
}

func (q *Question525) Run() {
	tests := []testcase525{
		{
			nums: []int{0, 1},
			ans:  2,
		},
		{
			nums: []int{0, 1, 0},
			ans:  2,
		},
		{ //-1, -2, -1, -2, -3, -4, -3, -2
			nums: []int{0, 0, 1, 0, 0, 0, 1, 1},
			ans:  6,
		},
	}

	for i, t := range tests {
		res := solution525(t.nums)
		if res != t.ans {
			log.Fatalf("[Question525] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question525] All tests passed")
}
