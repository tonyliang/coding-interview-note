package offer05

import (
	"fmt"
	"log"
)

func replaceSpace(s string) string {
	result := ""
	for _, c := range s {
		if rune(c) == rune(' ') {
			result += "%20"
		} else {
			result += string(c)
		}
	}
	return result
}

type TestCase struct {
	s   string
	ans string
}

func Run() {
	tests := []TestCase{
		{
			s:   " ",
			ans: "%20",
		},
		{
			s:   "a b",
			ans: "a%20b",
		},
		{
			s:   "",
			ans: "",
		},
		{
			s:   "  ok  ",
			ans: "%20%20ok%20%20",
		},
	}
	for i, t := range tests {
		res := replaceSpace(t.s)
		if res != t.ans {
			log.Fatalf("Test %d Failed. Expect %v, Got %v\n", i, t.ans, res)
		}
	}
	fmt.Println("All tests passed")
}
