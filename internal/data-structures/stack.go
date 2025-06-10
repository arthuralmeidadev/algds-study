package datastructures

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop(item T) T {
	if len(s.items) == 0 {
		return *new(T)
	}

	popped := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return popped
}
