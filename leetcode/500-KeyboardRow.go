package leetcode

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
)

type Question500 struct{}

type testcase500 struct {
	param []string
	ans   []string
}

func (q *Question500) Run() {
	tests := []testcase500{
		{
			param: []string{"Hello", "Alaska", "Dad", "Peace"},
			ans:   []string{"Alaska", "Dad"},
		},
		{
			param: []string{},
			ans:   []string{},
		},
		{
			param: []string{"zmvb", "qaz", "bb", "OO", "OoOoP"},
			ans:   []string{"zmvb", "bb", "OO", "OoOoP"},
		},
	}
	for i, t := range tests {
		result := solution500(t.param)
		if !helper.Equal(result, t.ans) {
			log.Fatalf("Question500 %dth test failed. Expect %v Got %v\n", i, t.ans, result)
		}
	}
	fmt.Println("[Question500], all tests passed")
}
