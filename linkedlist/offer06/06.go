//https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
package offer06

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/linkedlist/util"
)

type Result struct {
	data []int
}

//This implementation consumes less space but needs to iterate the list twice.
func reversePrint(head *util.ListNode) []int {
	if head == nil {
		return make([]int, 0)
	}
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	result := make([]int, count)
	cur = head
	for i := count - 1; i >= 0; i-- {
		result[i] = cur.Val
		cur = cur.Next
	}
	return result
}

//This implementation consumes more space.
func reversePrint2(head *util.ListNode) []int {
	result := &Result{
		data: make([]int, 0),
	}
	if head == nil {
		return make([]int, 0)
	}
	helper(head, result)
	return result.data
}

func helper(head *util.ListNode, r *Result) {
	if head.Next != nil {
		helper(head.Next, r)
	}
	r.data = append(r.data, head.Val)
}

type TestCase struct {
	head *util.ListNode
	ans  []int
}

func Run() {
	tests := []TestCase{
		{
			head: util.NewLinkedListFromSlice([]int{1, 2, 3}),
			ans:  []int{3, 2, 1},
		},
		{
			head: nil,
			ans:  []int{},
		},
		{
			head: util.NewLinkedListFromSlice([]int{3, 4, 5, 6}),
			ans:  []int{6, 5, 4, 3},
		},
	}
	for i, t := range tests {
		res := reversePrint(t.head)
		if !util.Equal(res, t.ans) {
			log.Fatalf("Test %d Failed. Expect %v, Got %v", i, t.ans, res)
		}
	}
	fmt.Println("All tests passed")
}
