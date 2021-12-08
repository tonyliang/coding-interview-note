package leetcode

import (
	"fmt"
	"log"
)

type Question581 struct{}
type testcase581 struct {
	nums []int
	ans  int
}

func (q *Question581) Run() {
	tests := []testcase581{
		{
			nums: []int{2, 6, 4, 8, 10, 9, 15},
			ans:  5,
		},
		{
			nums: []int{1, 2, 3, 4},
			ans:  0,
		},
		{
			nums: []int{1},
			ans:  0,
		},
	}
	for i, t := range tests {
		res := solution581(t.nums)
		if res != t.ans {
			log.Fatalf("[Question 581] test %d failed, expect: %v, got: %v\n", i, t.ans, res)
		}
	}
	fmt.Println("[Question 581] All tests passed")
}
