package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

// Insert new node
func (bst *BinarySearchTree) InsertRec(node *Node, val int) *Node {
	if bst.root == nil {
		bst.root = &Node{val, nil, nil}
		return bst.root
	}

	if node == nil {
		return &Node{val, nil, nil}
	}

	if val <= node.data {
		node.left = bst.InsertRec(node.left, val)
	} else {
		node.right = bst.InsertRec(node.right, val)
	}

	return node
}

func (bst *BinarySearchTree) Insert(val int) {
	bst.InsertRec(bst.root, val)
}

// Search node
func (bst *BinarySearchTree) SearchRec(node *Node, val int) bool {
	if node == nil {
		return false
	}

	if node.data == val {
		return true
	}

	if val < node.data {
		return bst.SearchRec(node.left, val)
	} else {
		return bst.SearchRec(node.right, val)
	}
}

func (bst *BinarySearchTree) Search(val int) bool {
	found := bst.SearchRec(bst.root, val)
	return found
}

// Traversal
func (bst *BinarySearchTree) InOrder(node *Node) {
	if node == nil {
		return
	} else {
		bst.InOrder(node.left)
		fmt.Print(node.data, " ")
		bst.InOrder(node.right)
	}
}

func (bst *BinarySearchTree) PreOrder(node *Node) {
	if node == nil {
		return
	} else {
		fmt.Print(node.data, " ")
		bst.PreOrder(node.left)
		bst.PreOrder(node.right)
	}
}

func (bst *BinarySearchTree) PostOrder(node *Node) {
	if node == nil {
		return
	} else {
		bst.PostOrder(node.left)
		bst.PostOrder(node.right)
		fmt.Print(node.data, " ")
	}
}

func main() {
	bst := BinarySearchTree{}
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(20)
	bst.Insert(17)
	bst.Insert(4)
	bst.Insert(6)
	fmt.Println()
	bst.InOrder(bst.root)
	fmt.Println()
	bst.PreOrder(bst.root)
	fmt.Println()
	fmt.Println(bst.Search(5))
	fmt.Println()
	bst.PostOrder(bst.root)
}
