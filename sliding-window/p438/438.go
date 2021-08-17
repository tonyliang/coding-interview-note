package p438

import (
	"fmt"
	"log"
)

func findAnagrams(s string, p string) []int {
	if len(s) == 0 || len(p) == 0 {
		return make([]int, 0)
	}
	result := make([]int, 0)
	pmap := make(map[rune]int)

	//build pattern map
	for _, c := range p {
		pmap[c] += 1
	}

	left := 0
	right := 0
	matchCount := 0
	curmap := make(map[rune]int)

	for right < len(s) {
		if v, ok := pmap[rune(s[right])]; ok {
			curmap[rune(s[right])]++
			if curmap[rune(s[right])] == v {
				matchCount++
			} else if curmap[rune(s[right])] == v+1 {
				matchCount--
			}
			for curmap[rune(s[right])] > v {
				curmap[rune(s[left])]--
				if curmap[rune(s[left])] == pmap[rune(s[left])] {
					matchCount++
				} else if curmap[rune(s[left])] == pmap[rune(s[left])]-1 {
					matchCount--
				}
				left++
			}
			if matchCount == len(pmap) {
				result = append(result, left)
			}
		} else {
			matchCount = 0
			curmap = make(map[rune]int)
			left = right + 1
		}
		right++
	}
	return result
}

type TestCase struct {
	s   string
	p   string
	ans []int
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Run() {
	tests := []TestCase{
		{
			s:   "",
			p:   "",
			ans: []int{},
		},
		{
			s:   "ab",
			p:   "ab",
			ans: []int{0},
		},
		{
			s:   "abcabc",
			p:   "ab",
			ans: []int{0, 3},
		},
		{
			s:   "ab",
			p:   "ab",
			ans: []int{0},
		},
		{
			s:   "cbaebabacd",
			p:   "abc",
			ans: []int{0, 6},
		},
		{
			s:   "abab",
			p:   "ab",
			ans: []int{0, 1, 2},
		},
	}
	for i, t := range tests {
		res := findAnagrams(t.s, t.p)
		if !equal(res, t.ans) {
			log.Fatalf("Test %d failed. Expect: %v, Got: %v", i, t.ans, res)
		}
	}
	fmt.Println("All tests passed")
}
