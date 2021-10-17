package leetcode

import (
	"fmt"
	"log"
)

type Question523 struct{}

type testcase523 struct {
	nums []int
	k    int
	ans  bool
}

func (q *Question523) Run() {
	tests := []testcase523{
		{
			nums: []int{23, 2, 4, 6, 7},
			k:    6,
			ans:  true,
		},
		{
			nums: []int{23, 2, 6, 4, 7},
			k:    6,
			ans:  true,
		},
		{
			nums: []int{23, 2, 4, 6, 7},
			k:    13,
			ans:  true,
		},
		{
			nums: []int{23, 2, 6, 4, 7},
			k:    13,
			ans:  false,
		},
		//this case is why `seen[0] = -1` is used.
		{
			nums: []int{23, 2, 4, 6, 6},
			k:    7,
			ans:  true,
		},
		{
			nums: []int{6},
			k:    6,
			ans:  false,
		},
		{
			nums: []int{1, 2, 3},
			k:    6,
			ans:  true,
		},
	}
	for i, t := range tests {
		res := solution523(t.nums, t.k)
		if res != t.ans {
			log.Fatalf("[Question523] test %d failed, expect %v got %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question523] All tests passed")
}
