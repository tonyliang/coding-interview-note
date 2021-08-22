package tree

import (
	"fmt"
	"log"
)

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	//find all matching root node of B in A
	q := []*TreeNode{A}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		if t.Left != nil {
			q = append(q, t.Left)
		}
		if t.Right != nil {
			q = append(q, t.Right)
		}
		if t.Val == B.Val {
			if isSubStructureHelper(t, B) {
				return true
			}
		}
	}
	return false
}

func isSubStructureHelper(x *TreeNode, y *TreeNode) bool {
	if y == nil {
		return true
	}
	if x == nil {
		return false
	}
	if x.Val == y.Val {
		return isSubStructureHelper(x.Left, y.Left) && isSubStructureHelper(x.Right, y.Right)
	}
	return false
}

type TestCaseOffer26 struct {
	A   *TreeNode
	B   *TreeNode
	ans bool
}

type Offer26 struct{}

func (o *Offer26) Run() {
	tests := []TestCaseOffer26{
		{
			A:   buildTreeFromString("3,4,5,1,2,null,null,null,null,null,null"),
			B:   buildTreeFromString("4,1,null,null,null"),
			ans: true,
		},
		{
			A:   buildTreeFromString("1,2,3,null,null,null,null"),
			B:   buildTreeFromString("3,1,null,null,null"),
			ans: false,
		},
		{
			A:   buildTreeFromString("1,4,3,null,null,5,7,null,null,null,null"),
			B:   buildTreeFromString("3,null,5,null,7,null,null"),
			ans: false,
		},
	}
	for i, t := range tests {
		res := isSubStructure(t.A, t.B)
		if res != t.ans {
			log.Fatalf("[Offer26] Test %d Failed. Expect %v, Got %v", i, t.ans, res)
		}
	}
	fmt.Println("[Offer26] All tests passed")
}
