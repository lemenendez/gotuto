package main

import (
	"fmt"
)

type BinaryNode struct {
	left *BinaryNode
	right *BinaryNode
	data rune
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert (data rune) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryTree{data:data, left:nil, right:nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func main() {
	tree := &BinaryTree{}

}