package leetcode

import (
	"fmt"
	"log"
	"sort"

	"github.com/tonyliang/coding-interview-note/helper"
)

type Question508 struct{}

type testcase508 struct {
	nums []int
	ans  []int
}

func (q *Question508) Run() {
	tests := []testcase508{
		{
			nums: []int{5, 2, -3},
			ans:  []int{2, -3, 4},
		},
		{
			nums: []int{5, 2, -5},
			ans:  []int{2},
		},
	}
	for i, t := range tests {
		root := BuildTree(t.nums)
		res := solution508(root)
		sort.Slice(t.ans, func(i, j int) bool {
			return t.ans[i] < t.ans[j]
		})
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		if !helper.Equal(res, t.ans) {
			log.Fatalf("[Question508] test %d failed, expect: %v got %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question508] All tests passed")
}
