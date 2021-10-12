package leetcode

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
)

type Question496 struct{}
type testcase496 struct {
	nums1 []int
	nums2 []int
	ans   []int
}

func (q *Question496) Run() {
	tests := []testcase496{
		{
			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2},
			ans:   []int{-1, 3, -1},
		},
		{
			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			ans:   []int{3, -1},
		},
	}
	for i, t := range tests {
		res := solution496(t.nums1, t.nums2)
		if !helper.Equal(res, t.ans) {
			log.Fatalf("[Question496] test %d failed, expect %v got %v\n", i, t.ans, res)
		}
	}
	fmt.Println("[Question496] All tests passed")
}
