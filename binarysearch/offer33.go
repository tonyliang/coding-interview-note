package binarysearch

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
)

/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。



参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true


提示：

数组长度 <= 1000


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//O(NlogN)
func verifyPostorderX(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	var dfs func([]int, int, int) bool
	dfs = func(arr []int, min, max int) bool {
		if len(arr) == 0 {
			return true
		}
		root := arr[len(arr)-1]
		if root < min || root > max {
			return false
		}
		i := 0
		for i < len(arr)-1 && arr[i] < root {
			i++
		}
		left := arr[0:i]
		right := arr[i : len(arr)-1]

		return dfs(left, min, root) && dfs(right, root, max)
	}
	return dfs(postorder, helper.MinInt, helper.MaxInt)
}

/* smarter: O(N) */
func verifyPostorder(postorder []int) bool {
	//initialize root node to be a super big number
	root := 1 << (32 - 1)

	stack := make([]int, 0)

	for i := len(postorder) - 1; i >= 0; i-- {

		//[check left nodes]
		//any node (left node) who's greater than its root node, break the tree structure
		//hence return false immediately.
		if postorder[i] > root {
			return false
		}

		//[check root node]
		//this loop will find the root node for current postorder[i] node.
		for len(stack) > 0 && stack[len(stack)-1] > postorder[i] {
			root = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
		}

		//[check right nodes]
		//stack is monotonically increasing
		//storing larger (right child) nodes
		stack = append(stack, postorder[i])
	}
	return true
}

type Offer33 struct{}
type TestCaseOffer33 struct {
	postorder []int
	ans       bool
}

func (o *Offer33) Run() {
	tests := []TestCaseOffer33{
		{
			postorder: []int{1, 6, 3, 2, 5},
			ans:       false,
		},
		{
			postorder: []int{1, 3, 2, 6, 5},
			ans:       true,
		},
		{
			postorder: []int{1},
			ans:       true,
		},
	}
	for i, t := range tests {
		res := verifyPostorder(t.postorder)
		if res != t.ans {
			log.Fatalf("[Offer33] test %d failed, expect: %v, got: %v", i, t.ans, res)
		}
	}
	fmt.Println("[Offer33] All tests passed")
}
