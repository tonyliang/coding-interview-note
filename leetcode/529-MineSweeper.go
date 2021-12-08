package leetcode

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
)

type Question529 struct{}
type testcase529 struct {
	board [][]byte
	click []int
	ans   [][]byte
}

func (q *Question529) Run() {
	tests := []testcase529{
		{
			board: [][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			},
			click: []int{3, 0},
			ans: [][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
		},
		{
			board: [][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'M', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
			click: []int{1, 2},
			ans: [][]byte{
				{'B', '1', 'E', '1', 'B'},
				{'B', '1', 'X', '1', 'B'},
				{'B', '1', '1', '1', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
		},
	}
	for i, t := range tests {
		res := solution529(t.board, t.click)
		if !helper.Equal(res, t.ans) {
			log.Fatalf("[Question529] test %d failed, expect: %v, got: %v\n", i, t.ans, res)
		}
	}
	fmt.Println("[Question529] All tests passed")
}
