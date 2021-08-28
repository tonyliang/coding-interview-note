package array

import (
	"fmt"
	"log"

	h "github.com/tonyliang/coding-interview-note/helper"
)

/*
 Given two array of the same size, A and B.
 One array represents the correct index for each element in the other array.
 For example:
   A = [2,4,1,3,0]
   B = ["tony", "jack", "lynn", "mark", "mike"]

 Write a function to return sorted Array (in ascending order) of B based on A.

  Example answer: C = ["mike", "lynn", "tony", "mark", "jack"]

  NOTE: the indice in A is between 0 ~ N-1, where N is the length of A
  NOTE: using a hashmap is trivial, can you achieve it with O(1) space?
*/

func Sorted(order []int, value []string) []string {
	//Solution is to swap each element in A until each element is at its right location.
	for i := 0; i < len(order); {
		if order[i] == i {
			i++
		} else {
			swapHelper(order, value, i, order[i])
		}
	}
	return value
}

func swapHelper(A []int, B []string, i, j int) {
	t := A[i]
	A[i] = A[j]
	A[j] = t

	x := B[i]
	B[i] = B[j]
	B[j] = x
}

type X1 struct{}
type TestCaseX1 struct {
	A   []int
	B   []string
	ans []string
}

func (x *X1) Run() {
	tests := []TestCaseX1{
		{
			A:   []int{2, 4, 1, 3, 0},
			B:   []string{"tony", "jack", "lynn", "mark", "mike"},
			ans: []string{"mike", "lynn", "tony", "mark", "jack"},
		},
		{
			A:   []int{0, 1, 2},
			B:   []string{"ja", "xy", "zh"},
			ans: []string{"ja", "xy", "zh"},
		},
		{
			A:   []int{2, 1, 0},
			B:   []string{"wo", "eee", "fine"},
			ans: []string{"fine", "eee", "wo"},
		},
	}
	for i, t := range tests {
		res := Sorted(t.A, t.B)
		if !h.Equal(t.ans, res) {
			log.Fatalf("[X1] Test %d Failed. Expect %v, Got %v", i, t.ans, res)
		}
	}
	fmt.Println("[X1] All tests passed")
}
