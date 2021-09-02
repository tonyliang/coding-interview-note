package tree

import (
	"fmt"
	"log"

	"github.com/tonyliang/coding-interview-note/helper"
)

/*
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。



示例:
给定如下二叉树，以及目标和 target = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]


提示：

节点总数 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	var helper func(*TreeNode, []int, int)
	helper = func(node *TreeNode, path []int, total int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if total+node.Val == target {
				path = append(path, node.Val)
				dest := make([]int, len(path))
				copy(dest, path)
				ans = append(ans, dest)
			}
		} else {
			path = append(path, node.Val)
			helper(node.Left, path, total+node.Val)
			helper(node.Right, path, total+node.Val)
			path = path[0 : len(path)-1]
		}
	}

	helper(root, make([]int, 0), 0)
	return ans
}

//a more elegant one from someone else
/*
func pathSum(root *TreeNode, target int) (paths [][]int) {
    path := []int{}
    var dfs func(root *TreeNode, left int)
    dfs = func(root *TreeNode, left int) {
        if root == nil {
            return
        }

        left -= root.Val
        path = append(path, root.Val)
        defer func() {path = path[:len(path)-1]}()
        if root.Left == nil && root.Right == nil && left == 0 {
            paths = append(paths, append([]int(nil), path...))
        }

        dfs(root.Left, left)
        dfs(root.Right, left)
    }

    dfs(root, target)
    return
}
*/
type Offer34 struct{}
type TestCaseOffer34 struct {
	root   *TreeNode
	target int
	ans    [][]int
}

func (o *Offer34) Run() {
	tests := []TestCaseOffer34{
		{
			root:   BuildTreeFromString("5,4,8,11,null,13,4,7,2,null,null,5,1,null,null,null,null,null,null,null,null"),
			target: 22,
			ans:    [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
	}
	for i, t := range tests {
		res := pathSum(t.root, t.target)
		fmt.Println(res)
		for j, r := range res {
			if !helper.Equal(r, t.ans[j]) {
				log.Fatalf("[Offer34] test %d failed, expect: %v, got: %v", i, t.ans[j], r)
			}
		}
	}
	fmt.Println("[Offer34] All tests passed")
}
