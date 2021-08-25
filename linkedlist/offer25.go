package linkedlist

import (
	"fmt"
	"log"
)

/*
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：

0 <= 链表长度 <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	p1 := l1
	p2 := l2
	dummy := &ListNode{}
	p3 := dummy
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p3.Next = p1
			p1 = p1.Next
		} else {
			p3.Next = p2
			p2 = p2.Next
		}
		p3 = p3.Next
	}
	if p1 != nil {
		p3.Next = p1
	}
	if p2 != nil {
		p3.Next = p2
	}
	return dummy.Next
}

type Offer25 struct{}

type TestCaseOffer25 struct {
	l1  *ListNode
	l2  *ListNode
	ans *ListNode
}

func (o *Offer25) Run() {
	tests := []TestCaseOffer25{
		{
			l1:  NewLinkedListFromSlice([]int{1, 2, 4}),
			l2:  NewLinkedListFromSlice([]int{1, 3, 4}),
			ans: NewLinkedListFromSlice([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			l1:  NewLinkedListFromSlice([]int{1, 2, 3}),
			l2:  NewLinkedListFromSlice([]int{2}),
			ans: NewLinkedListFromSlice([]int{1, 2, 2, 3}),
		},
		{
			l1:  nil,
			l2:  NewLinkedListFromSlice([]int{1, 2}),
			ans: NewLinkedListFromSlice([]int{1, 2}),
		},
	}
	for i, t := range tests {
		res := mergeTwoLists(t.l1, t.l2)
		if !EqualLinkedList(t.ans, res) {
			log.Fatalf("[Offer25] test %d failed, expect %v, got %v\n", i, t.ans, res)
		}
	}
	fmt.Println("[Offer25] All tests passed")
}
