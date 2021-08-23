package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedListFromSlice(data []int) *ListNode {
	if len(data) == 0 {
		return nil
	}
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := dummy
	for i := 0; i < len(data); i++ {
		cur.Next = &ListNode{
			Val:  data[i],
			Next: nil,
		}
		cur = cur.Next
	}
	return dummy.Next
}

func EqualLinkedList(a *ListNode, b *ListNode) bool {
	ca := a
	cb := b
	for {
		if ca == nil && cb == nil {
			break
		}
		if ca == nil || cb == nil {
			return false
		}
		if ca.Val != cb.Val {
			return false
		}
		ca = ca.Next
		cb = cb.Next
	}
	return true
}

func PrintLinkedList(head *ListNode) []int {
	result := make([]int, 0)
	cur := head
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return result
}

func Equal2(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !Equal(v, b[i]) {
			return false
		}
	}
	return true
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
