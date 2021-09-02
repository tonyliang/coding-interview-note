package tree

import (
	"fmt"
	"log"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	idx := indexOf(inorder, preorder[0])

	sizeOfLeft := idx
	root.Left = buildTree(preorder[1:1+sizeOfLeft], inorder[0:idx])
	root.Right = buildTree(preorder[1+sizeOfLeft:], inorder[idx+1:])
	return root
}

func indexOf(haystack []int, needle int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

type Offer07 struct {
}

type TestCaseOffer07 struct {
	preorder []int
	inorder  []int
	ans      *TreeNode
}

func (o *Offer07) Run() {
	tests := []TestCaseOffer07{
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			ans:      BuildTreeFromString("3,9,20,null,null,15,7,null,null,null,null"),
		},
		{
			preorder: []int{-1},
			inorder:  []int{-1},
			ans:      &TreeNode{Val: -1},
		},
	}
	for i, t := range tests {
		res := buildTree(t.preorder, t.inorder)
		resPreorder := preOrderString(res)
		resInorder := getInorderString(res)
		ansPreorder := preOrderString(t.ans)
		ansInorder := getInorderString(t.ans)
		if resPreorder != ansPreorder || resInorder != ansInorder {
			log.Fatalf("[Offer07] test %d failed, expect: %v, got: %v", i, ansPreorder, resPreorder)
		}
	}
	fmt.Println("[Offer07] All tests passed")
}
