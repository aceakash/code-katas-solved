package balanced

import (
	"go-katas/stack"
)

func IsBalanced(expr string) bool {
	capacity := len(expr)/2 + 1
	s := stack.New[rune](capacity)

	for _, r := range expr {
		if !isOpener(r) && !isCloser(r) {
			return false
		}
		if isOpener(r) {
			s.Push(r)
			continue
		}
		// is closer
		popped, couldPop := s.Pop()
		if popped != openerFor(r) || !couldPop {
			return false
		}
	}
	return s.IsEmpty()
}

func isOpener(r rune) bool {
	return r == '(' || r == '{' || r == '['
}
func isCloser(r rune) bool {
	return r == ')' || r == '}' || r == ']'
}

func openerFor(r rune) rune {
	switch r {
	case ')':
		return '('
	case '}':
		return '{'
	case ']':
		return '['
	}

	return 0
}
