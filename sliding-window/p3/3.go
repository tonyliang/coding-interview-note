package p3

import (
	"fmt"
	"log"
)

func lengthOfLongestSubstring(s string) int {
	ans := 0
	seen := make(map[byte]int)
	left, right := 0, 0
	for right < len(s) {
		seen[s[right]]++

		for seen[s[right]] > 1 {
			seen[s[left]]--
			left++
		}

		ans = Max(ans, right-left+1)
		right++
	}
	return ans
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type TestCase struct {
	s   string
	ans int
}

func Run() {
	tests := []TestCase{
		{
			s:   "",
			ans: 0,
		},
		{
			s:   "a",
			ans: 1,
		},
		{
			s:   "aaa",
			ans: 1,
		},
		{
			s:   "abababcaaa",
			ans: 3,
		},
		{
			s:   "abcabc",
			ans: 3,
		},
		{
			s:   "abcabcbb",
			ans: 3,
		},
		{
			s:   "bbbbb",
			ans: 1,
		},
		{
			s:   "pwwkew",
			ans: 3,
		},
	}
	for i, t := range tests {
		res := lengthOfLongestSubstring(t.s)
		if res != t.ans {
			log.Fatalf("test %d failed. Expect %d, Got: %d\n", i, t.ans, res)
		}
	}
	fmt.Println("All tests passed")
}
