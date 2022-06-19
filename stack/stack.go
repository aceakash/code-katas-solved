package stack

type Stack[T any] struct {
	storage []T
}

func New[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		storage: make([]T, 0, capacity),
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.storage) == 0
}

func (s *Stack[T]) Push(item T) bool {
	if s.IsFull() {
		return false
	}
	s.storage = append(s.storage, item)
	return true
}

func (s *Stack[T]) Pop() (T, bool) {
	len := len(s.storage)
	if len == 0 {
		var zero T
		return zero, false
	}

	popped := s.storage[len-1]
	s.storage = s.storage[0 : len-1]
	return popped, true
}

func (s *Stack[T]) IsFull() bool {
	if cap(s.storage) == len(s.storage) {
		return true
	}
	return false
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.storage[len(s.storage)-1], true
}
