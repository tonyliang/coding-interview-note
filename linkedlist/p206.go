package linkedlist

import (
	"fmt"
	"log"
)

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

type P206 struct{}

type TestCaseP206 struct {
	head *ListNode
	ans  *ListNode
}

func (p *P206) Run() {
	tests := []TestCaseP206{
		{
			head: NewLinkedListFromSlice([]int{1, 2, 3, 4, 5}),
			ans:  NewLinkedListFromSlice([]int{5, 4, 3, 2, 1}),
		},
	}
	for i, t := range tests {
		res := reverseList(t.head)
		if !EqualLinkedList(res, t.ans) {
			log.Fatalf("[p206] Test %d Failed. Expect %d, Got %d", i, PrintLinkedList(t.ans), PrintLinkedList(res))
		}
		fmt.Println("[p206] All tests passed")
	}
}
