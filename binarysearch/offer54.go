package binarysearch

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/tree"
)

/*
给定一棵二叉搜索树，请找出其中第k大的节点。



示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 4
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func kthLargest(root *tree.TreeNode, k int) int {
	cur := 0
	var ans *tree.TreeNode
	var helper func(node *tree.TreeNode)
	helper = func(node *tree.TreeNode) {
		if node != nil {
			helper(node.Right)
			cur++
			if cur == k {
				ans = node
				return
			}
			helper(node.Left)
		}
	}
	helper(root)
	return ans.Val
}

type Offer54 struct{}
type TestCaseOffer54 struct {
	root *tree.TreeNode
	k    int
	ans  int
}

func (o *Offer54) Run() {
	tests := []TestCaseOffer54{
		{
			root: tree.BuildTreeFromString("3,1,4,null,2,null,null,null,null"),
			k:    1,
			ans:  4,
		},
		{
			root: tree.BuildTreeFromString("5,3,6,2,4,null,null,1,null,null,null,null,null"),
			k:    3,
			ans:  4,
		},
	}
	for i, t := range tests {
		res := kthLargest(t.root, t.k)
		if res != t.ans {
			log.Fatalf("[Offer54] test %d failed, expect %v, got %v", i, t.ans, res)
		}
	}
	fmt.Println("[Offer54] All tests passed")
}
