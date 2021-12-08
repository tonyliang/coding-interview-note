package helper

import (
	"fmt"
	"log"
)

type BSTNode struct {
	Key   int
	Count int
	Left  *BSTNode
	Right *BSTNode
}

func NewBSTNode(key int) *BSTNode {
	node := BSTNode{Key: key, Count: 1}
	return &node
}

func BSTInsert(node *BSTNode, key int) *BSTNode {
	if node == nil {
		return NewBSTNode(key)
	}
	if node.Key < key {
		node.Right = BSTInsert(node.Right, key)
	}
	if node.Key > key {
		node.Left = BSTInsert(node.Left, key)
	}
	if node.Key == key {
		node.Count++
	}

	return node
}

func BSTDelete(node *BSTNode, key int) *BSTNode {
	if node == nil {
		return nil
	}
	if key > node.Key {
		node.Right = BSTDelete(node.Right, key)
	} else if key < node.Key {
		node.Left = BSTDelete(node.Left, key)
	} else {
		node.Count--
		if node.Count == 0 {
			if node.Right == nil && node.Left == nil {
				return nil
			}
			if node.Right == nil {
				return node.Left
			}
			if node.Left == nil {
				return node.Right
			}
			//both right and left aren't nil
			//find the min node of right subtree
			minNode := node

			for minNode.Left != nil {
				minNode = minNode.Left
			}

			node.Key = minNode.Key
			node.Count = minNode.Count
			node.Right = BSTDelete(node.Right, minNode.Key)
			return node

		}
	}
	return node
}

func BSTSearchMax(node *BSTNode) *BSTNode {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		return BSTSearchMax(node.Right)
	}
	return node
}

func TestBST() {
	var root *BSTNode
	root = BSTInsert(root, 12)
	root = BSTInsert(root, 10)
	root = BSTInsert(root, 20)
	root = BSTInsert(root, 9)
	root = BSTInsert(root, 11)
	root = BSTInsert(root, 10)
	root = BSTInsert(root, 12)
	root = BSTInsert(root, 12)
	result := BSTSearchMax(root)

	if result == nil {
		log.Fatal("error finding max in BST")
	}
	if result.Count != 1 {
		log.Fatalf("expect %d got: %d", 1, result.Count)
	}
	if result.Key != 20 {
		log.Fatalf("expect %d got: %d", 20, result.Key)
	}
	root = BSTDelete(root, 20)
	result = BSTSearchMax(root)
	if result.Count != 3 {
		log.Fatalf("expect %d got: %d", 3, result.Count)
	}
	if result.Key != 12 {
		log.Fatalf("expect %d got: %d", 12, result.Key)
	}

	for i := 0; i < 2; i++ {
		root = BSTDelete(root, 12)
		result = BSTSearchMax(root)
		if result.Count != 2-i {
			log.Fatalf("%d, expect %d got: %d", i, 2-i, result.Count)
		}
		if result.Key != 12 {
			log.Fatalf("%d, expect %d got: %d", i, 12, result.Key)
		}
	}
	root = BSTDelete(root, 12)
	result = BSTSearchMax(root)
	if result.Count != 1 {
		log.Fatalf("expect %d got: %d", 1, result.Count)
	}
	if result.Key != 11 {
		log.Fatalf("expect %d got: %d", 11, result.Key)
	}
	fmt.Println("BSTNode Tests Passed")
}
