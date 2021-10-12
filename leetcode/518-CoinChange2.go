package leetcode

import (
	"fmt"
	"log"
)

type Question518 struct{}

type testcase518 struct {
	amount int
	coins  []int
	ans    int
}

func (q *Question518) Run() {
	tests := []testcase518{
		{
			amount: 5,
			coins:  []int{1, 2, 5},
			ans:    4,
		},
		{
			amount: 0,
			coins:  []int{7},
			ans:    1,
		},
		{
			amount: 3,
			coins:  []int{1, 2, 5},
			ans:    2,
		},
		{
			amount: 3,
			coins:  []int{2},
			ans:    0,
		},
		{
			amount: 10,
			coins:  []int{10},
			ans:    1,
		},
	}
	for i, t := range tests {
		res := solution518(t.amount, t.coins)
		if res != t.ans {
			log.Fatalf("[Question518] test %d failed, expect %v got %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question518] All tests passed")
}
