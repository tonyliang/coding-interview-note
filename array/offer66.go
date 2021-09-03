package array

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
)

func constructArr(a []int) []int {
	/*
	   given [a,b,c,d]
	   want [bcd,acd,abd,abc]
	   think  [1,a,b,c] =>  [1,    a,  ab,  abc]
	   think  [b,c,d,1] =>  [bcd,  cd,  d,     1]
	*/
	leftToRight := make([]int, len(a))
	rightToLeft := make([]int, len(a))
	ans := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		if i == 0 {
			leftToRight[i] = 1
		} else {
			leftToRight[i] = leftToRight[i-1] * a[i-1]
		}
	}

	for i := len(a) - 1; i >= 0; i-- {
		if i == len(a)-1 {
			rightToLeft[i] = 1
		} else {
			rightToLeft[i] = rightToLeft[i+1] * a[i+1]
		}
	}

	for i := 0; i < len(a); i++ {
		ans[i] = leftToRight[i] * rightToLeft[i]
	}

	return ans
}

type Offer66 struct{}
type TestCaseOffer66 struct {
	a   []int
	ans []int
}

func (o *Offer66) Run() {
	tests := []TestCaseOffer66{
		{
			a:   []int{1, 2, 3, 4, 5},
			ans: []int{120, 60, 40, 30, 24},
		},
		{
			a:   []int{5, 6},
			ans: []int{6, 5},
		},
		{
			a:   []int{5},
			ans: []int{1},
		},
	}
	for i, t := range tests {
		res := constructArr(t.a)
		if !helper.Equal(res, t.ans) {
			log.Fatalf("[Offer66] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Offer66] All tests passed")
}
