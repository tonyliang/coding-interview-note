package linkedlist

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tonyliang/coding-interview-note/tree"
)

/*
https://leetcode.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list/
Convert a Binary Search Tree to a sorted Circular Doubly-Linked List in place.

You can think of the left and right pointers as synonymous to the predecessor and successor pointers in a doubly-linked list. For a circular doubly linked list, the predecessor of the first element is the last element, and the successor of the last element is the first element.

We want to do the transformation in place. After the transformation, the left pointer of the tree node should point to its predecessor, and the right pointer should point to its successor. You should return the pointer to the smallest element of the linked list.
*/

var P426Prev *tree.TreeNode

func treeToDoublyList(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}
	dummy := &tree.TreeNode{}
	P426Prev = dummy
	treeToDoublyListHelper(root)
	//[note] now prev is pointing to tail, because prev is the last node visited.
	P426Prev.Right = dummy.Right //point tail to head
	dummy.Right.Left = P426Prev
	return dummy.Right
}

func treeToDoublyListHelper(root *tree.TreeNode) {
	if root == nil {
		return
	}
	treeToDoublyListHelper(root.Left)
	root.Left = P426Prev
	P426Prev.Right = root
	P426Prev = root
	treeToDoublyListHelper(root.Right)
}

type P426 struct{}

func (p *P426) Run() {
	root := tree.BuildTreeFromString("4,2,5,1,3,null,null,null,null,null,null")
	head := treeToDoublyList(root)

	leftToRight, rightToLeft := strconv.Itoa(head.Val), strconv.Itoa(head.Left.Val)

	cur := head.Right
	for cur.Val != head.Val {
		leftToRight += "," + strconv.Itoa(cur.Val)
		cur = cur.Right
	}
	tail := head.Left
	cur = tail.Left
	for cur.Val != tail.Val {
		rightToLeft += "," + strconv.Itoa(cur.Val)
		cur = cur.Left
	}
	if leftToRight != "1,2,3,4,5" {
		log.Fatalf("[P426] Failed, expect [1,2,3,4,5], got %v", leftToRight)
	}
	if rightToLeft != "5,4,3,2,1" {
		log.Fatalf("[P426] Failed, expect [5,4,3,2,1], got %v", rightToLeft)
	}
	fmt.Println("[P426] Test passed")
}
