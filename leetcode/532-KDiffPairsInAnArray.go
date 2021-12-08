package leetcode

import (
	"fmt"
	"log"
)

type Question532 struct{}
type testcase532 struct {
	nums []int
	k    int
	ans  int
}

func (q *Question532) Run() {
	tests := []testcase532{
		{
			nums: []int{3, 1, 4, 1, 5},
			k:    2,
			ans:  2,
		},
		{
			nums: []int{1, 2, 3, 4, 5},
			k:    1,
			ans:  4,
		},
		{
			nums: []int{1, 3, 1, 5, 4},
			k:    0,
			ans:  1,
		},
		{
			nums: []int{5, 5, 3, 3},
			k:    2,
			ans:  1,
		},
	}
	for i, t := range tests {
		res := solution532(t.nums, t.k)
		if res != t.ans {
			log.Fatalf("[Question532] test %d failed. expect %v, got %v\n", i, t.ans, res)
		}
	}
	fmt.Println("[Question532] All tests passed")
}
