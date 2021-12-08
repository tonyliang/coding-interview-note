package leetcode

import (
	"fmt"
	"log"
)

type Question567 struct{}
type testcase567 struct {
	s1  string
	s2  string
	ans bool
}

func (q *Question567) Run() {
	tests := []testcase567{
		{
			s1:  "ab",
			s2:  "abab",
			ans: true,
		},
		{
			s1:  "abc",
			s2:  "cbaebabacd",
			ans: true,
		},
		{
			s1:  "abc",
			s2:  "",
			ans: false,
		},
		{
			s1:  "abc",
			s2:  "abacbabc",
			ans: true,
		},
		{
			s1:  "ab",
			s2:  "eidboaoo",
			ans: false,
		},
	}
	for i, t := range tests {
		res := solution567(t.s1, t.s2)
		if res != t.ans {
			log.Fatalf("[Question 567] test %d failed, expect: %v, got: %v\n", i, t.ans, res)
		}
	}
	fmt.Println("[Question 567] All tests passed")
}
