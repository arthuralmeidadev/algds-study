package datastructures

import "fmt"

type BinaryTreeNode[T any] struct {
	parent *BinaryTreeNode[T]
	left   *BinaryTreeNode[T]
	right  *BinaryTreeNode[T]
	depth  int
	value  T
}

type BinaryTree[T any] struct {
	root *BinaryTreeNode[T]
}

func (n *BinaryTreeNode[T]) NewChildLeft(v T) (*BinaryTreeNode[T], error) {
	if n.left != nil {
		return nil, fmt.Errorf("This noe already has a left child node")
	}

	n.left = &BinaryTreeNode[T]{
		parent: n,
		depth:  n.depth + 1,
		value:  v,
	}

	return n.left, nil
}

func (n *BinaryTreeNode[T]) NewChildRight(v T) (*BinaryTreeNode[T], error) {
	if n.left != nil {
		return nil, fmt.Errorf("This noe already has a right child node")
	}

	n.right = &BinaryTreeNode[T]{
		parent: n,
		depth:  n.depth + 1,
		value:  v,
	}

	return n.right, nil
}

func (t *BinaryTree[T]) NewRoot(v T) (*BinaryTreeNode[T], error) {
	if t.root != nil {
		return nil, fmt.Errorf("This tree already has a root node")
	}

	t.root = &BinaryTreeNode[T]{
		depth: 0,
		value: v,
	}

	return t.root, nil
}
