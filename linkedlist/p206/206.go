package p206

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/linkedlist/util"
)

func reverseList(head *util.ListNode) *util.ListNode {
	var pre *util.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

type TestCase struct {
	head *util.ListNode
	ans  *util.ListNode
}

func Run() {
	tests := []TestCase{
		{
			head: util.NewLinkedListFromSlice([]int{1, 2, 3, 4, 5}),
			ans:  util.NewLinkedListFromSlice([]int{5, 4, 3, 2, 1}),
		},
	}
	for i, t := range tests {
		res := reverseList(t.head)
		if !util.EqualLinkedList(res, t.ans) {
			log.Fatalf("Test %d Failed. Expect %d, Got %d", i, util.PrintLinkedList(t.ans), util.PrintLinkedList(res))
		}
		fmt.Println("All tests passed")
	}
}
