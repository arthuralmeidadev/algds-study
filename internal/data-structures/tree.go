package datastructures

import "fmt"

type TreeNode[T any] struct {
	parent   *TreeNode[T]
	children []*TreeNode[T]
	depth    int
	value    T
}

type Tree[T any] struct {
	root *TreeNode[T]
}

func (n *TreeNode[T]) NewChild(v T) *TreeNode[T] {
	child := &TreeNode[T]{
		depth:  n.depth + 1,
		parent: n,
		value:  v,
	}

	n.children = append(n.children, child)
	return child
}

func (t *Tree[T]) NewRoot(v T) (*TreeNode[T], error) {
	if t.root != nil {
		return nil, fmt.Errorf("This tree already has a root node")
	}

	t.root = &TreeNode[T]{
		depth: 0,
		value: v,
	}

	return t.root, nil
}
