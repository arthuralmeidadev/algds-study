package datastructures

type Node[T any] struct {
	parent   *Node[T]
	children []*Node[T]
	depth    int
	value    T
}

type Tree[T any] struct {
	root *Node[T]
}

func (n *Node[T]) NewChild(v T) {
	n.children = append(n.children, &Node[T]{
		depth:  n.depth + 1,
		parent: n,
		value:  v,
	})
}

func (t *Tree[T]) Plant(v T) {
	if t.root == nil {
		t.root = &Node[T]{
			depth: 0,
			value: v,
		}
	}
}
