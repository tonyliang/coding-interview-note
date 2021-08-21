package tree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTreeFromString(data string) *TreeNode {
	tokens := strings.Split(data, ",")
	if len(tokens) == 0 {
		return nil
	}
	index := 1
	rootVal, _ := strconv.Atoi(tokens[0])
	root := &TreeNode{
		Val: rootVal,
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		left := tokens[index]
		index++
		right := tokens[index]
		index++
		if left != "null" {
			leftVal, _ := strconv.Atoi(left)
			t.Left = &TreeNode{Val: leftVal}
			q = append(q, t.Left)
		}
		if right != "null" {
			rightVal, _ := strconv.Atoi(right)
			t.Right = &TreeNode{Val: rightVal}
			q = append(q, t.Right)
		}
	}
	return root
}

func preOrderString(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := strconv.Itoa(root.Val)
	left := preOrderString(root.Left)
	right := preOrderString(root.Right)
	if len(left) > 0 {
		result += "," + left
	}
	if len(right) > 0 {
		result += "," + right
	}
	return result
}

func getInorderString(root *TreeNode) string {
	if root == nil {
		return ""
	}
	left := getInorderString(root.Left)
	right := getInorderString(root.Right)
	result := ""
	if len(left) > 0 {
		result = left + ","
	}
	result += strconv.Itoa(root.Val)
	if len(right) > 0 {
		result += "," + right
	}
	return result
}
