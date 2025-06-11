package datastructures

type Queue[T any] struct {
	capacity uint
	items    []T
}

func (q *Queue[T]) Enqueue(item T) {
	if len(q.items) == int(q.capacity) {
		return
	}

	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	if len(q.items) == 0 {
		return *new(T)
	}

	first := q.items[0]
	q.items = q.items[1:]

	return first
}

func (q *Queue[T]) Peek() T {
	if len(q.items) == 0 {
		return *new(T)
	}

	return q.items[0]
}

func (q *Queue[T]) Some(f func(T, int) bool) bool {
	for i := 0; i < len(q.items); i++ {
		if f(q.items[i], i) {
			return true
		}
	}

	return false
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Length() int {
	return len(q.items)
}

func (q *Queue[T]) Capacity() uint {
	return q.capacity
}
