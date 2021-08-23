package linkedlist

import (
	"fmt"
	"log"
)

/*
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

返回删除后的链表的头节点。

注意：此题对比原题有改动

示例 1:

输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
示例 2:

输入: head = [4,5,1,9], val = 1
输出: [4,5,9]
解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{}
	dummy.Next = head
	cur := head
	pre := dummy

	for cur != nil && cur.Val != val {
		pre = cur
		cur = cur.Next
	}
	if cur != nil {
		pre.Next = cur.Next
	}
	return dummy.Next
}

type Offer18 struct{}
type TestCaseOffer18 struct {
	head *ListNode
	val  int
	ans  *ListNode
}

func (o *Offer18) Run() {
	tests := []TestCaseOffer18{
		{
			head: NewLinkedListFromSlice([]int{4, 5, 1, 9}),
			val:  5,
			ans:  NewLinkedListFromSlice([]int{4, 1, 9}),
		},
		{
			head: NewLinkedListFromSlice([]int{4, 5, 1, 9}),
			val:  1,
			ans:  NewLinkedListFromSlice([]int{4, 5, 9}),
		},
		{
			head: NewLinkedListFromSlice([]int{2, 0, 1, 3}),
			val:  0,
			ans:  NewLinkedListFromSlice([]int{2, 1, 3}),
		},
	}
	for i, t := range tests {
		res := deleteNode(t.head, t.val)
		if !EqualLinkedList(res, t.ans) {
			log.Fatalf("[Offer18] test %d failed, expect: %v, got: %v\n", i, PrintLinkedList(t.ans), PrintLinkedList(res))
		}
	}
	fmt.Println("[Offer18] All tests passed")
}
