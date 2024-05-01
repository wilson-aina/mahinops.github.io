package main

import (
	"fmt"
)

type node struct {
	data  int
	left  *node
	right *node
}

type bst struct {
	root *node
}

func (t *bst) insertVal(val int) {
	newNode := &node{data: val}

	if t.root == nil {
		t.root = newNode
		return
	}

	insertNode(t.root, newNode)
}

func insertNode(root, newNode *node) {
	if newNode.data < root.data {
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode)
		}
	} else {
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode)
		}
	}

}

func (t *bst) inOrderTraversal() {
	inOrder(t.root)
	fmt.Println()
}

func inOrder(node *node) {
	if node != nil {
		inOrder(node.left)
		fmt.Printf("%d ", node.data)
		inOrder(node.right)
	}
}

func (t *bst) levelOrderTraversal() {
	if t.root == nil {
		return
	}
	queue := []*node{t.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", node.data)

		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

}

func main() {
	bst := bst{}
	keys := []int{50, 30, 20, 40, 70, 60, 80}

	for _, val := range keys {
		bst.insertVal(val)
	}

	bst.inOrderTraversal()

	bst.levelOrderTraversal()
}
