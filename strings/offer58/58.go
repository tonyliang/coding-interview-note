package offer58

import (
	"fmt"
	"log"
)

func reverseLeftWords(s string, n int) string {
	size := len(s)
	if size == 0 || n == 0 {
		return s
	}
	count := n % size
	return s[count:] + s[:count]
}

type TestCase struct {
	s   string
	n   int
	ans string
}

func Run() {
	tests := []TestCase{
		{
			s:   "",
			n:   10,
			ans: "",
		},
		{
			s:   "ab",
			n:   2,
			ans: "ab",
		},
		{
			s:   "abc",
			n:   1,
			ans: "bca",
		},
		{
			s:   "abcdefg",
			n:   2,
			ans: "cdefgab",
		},
		{
			s:   "lrloseumgh",
			n:   6,
			ans: "umghlrlose",
		},
	}
	for i, t := range tests {
		res := reverseLeftWords(t.s, t.n)
		if res != t.ans {
			log.Fatalf("Test %d failed. Expect: %v, Got: %v", i, t.ans, res)
		}
	}
	fmt.Println("All tests passed")
}
