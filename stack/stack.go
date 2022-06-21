package stack

type Stack[T any] struct {
	storage []T
	top     int
}

func New[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		storage: make([]T, capacity),
		top:     -1,
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack[T]) Push(item T) bool {
	if s.IsFull() {
		return false
	}
	s.top += 1
	s.storage[s.top] = item
	return true
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	popped := s.storage[s.top]
	s.top--
	return popped, true
}

func (s *Stack[T]) IsFull() bool {
	return s.top == len(s.storage)-1
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.storage[s.top], true
}
