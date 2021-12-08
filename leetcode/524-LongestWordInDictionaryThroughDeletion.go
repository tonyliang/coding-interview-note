package leetcode

import (
	"fmt"
	"log"
)

type Question524 struct{}

type testcase524 struct {
	s   string
	d   []string
	ans string
}

func (q *Question524) Run() {
	tests := []testcase524{
		{
			s:   "abpcplea",
			d:   []string{"ale", "apple", "monkey", "plea"},
			ans: "apple",
		},
		{
			s:   "abpcplea",
			d:   []string{"a", "b", "c"},
			ans: "a",
		},
		{
			s:   "abpcplea",
			d:   []string{"aaaaa", "b", "c"},
			ans: "b",
		},
		{
			s:   "bab",
			d:   []string{"ba", "ab", "a", "b"},
			ans: "ab",
		},
		{
			s:   "aewfafwafjlwajflwajflwafj",
			d:   []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"},
			ans: "ewaf",
		},
	}
	for i, t := range tests {
		res := solution524(t.s, t.d)
		if res != t.ans {
			log.Fatalf("[Question524] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Question524] All tests passes")
}
