package datastructures

import "fmt"

type LinkedNode[T any] struct {
	head  *LinkedNode[T]
	tail  *LinkedNode[T]
	value T
}

type LinkedList[T any] struct {
	head *LinkedNode[T]
	tail *LinkedNode[T]
}

func (l *LinkedList[T]) Append(v T) {
	newNode := &LinkedNode[T]{
		head:  l.tail,
		value: v,
	}

	if l.head == nil {
		l.head = newNode
		return
	}

	if l.tail != nil {
		l.tail.tail = newNode
	}

	l.tail = newNode
}

func (l *LinkedList[T]) Prepend(v T) {
	newNode := &LinkedNode[T]{
		tail:  l.head,
		value: v,
	}

	if l.head != nil {
		l.head.head = newNode
		l.tail = l.head
	}

	l.head = newNode
}

func (l *LinkedList[T]) Insert(v T, index int) error {
	if index == 0 {
		l.Prepend(v)
		return nil
	}

	current := l.head

	for i := 0; i <= index; i++ {
		if current == l.tail {
			if i+1 != index {
				return fmt.Errorf("No links can be made at position %d", index)
			}

			l.Append(v)
			return nil
		}

		if i == index {
			if current.head == nil {
				return fmt.Errorf("No links can be made at position %d", index)
			}

			newNode := &LinkedNode[T]{
				head:  current.head,
				tail:  current,
				value: v,
			}

			current.head.tail = newNode

			return nil
		}

		current = current.tail
	}

	return nil
}

func (l *LinkedList[T]) FindFunc(f func(v T) bool) *T {
	current := l.head

	for {
		if current == nil {
			return nil
		}

		if f(current.value) {
			return &current.value
		}

		current = current.tail
	}
}

func (l *LinkedList[T]) DeleteFunc(f func(v T) bool) {
	current := l.head

	for {
		if current == nil {
			return
		}

		if f(current.value) {
			if current.head != nil && current.tail != nil {
				newHead, newTail := current.head, current.tail
				current.tail.head, current.head.tail = newTail, newHead
				return
			}

			if current.head != nil {
				current.head.tail = nil
			}

			if current.tail != nil {
				current.tail.head = nil
			}

			return
		}

		current = current.tail
	}
}
