package balanced

import (
	"go-katas/stack"
)

func IsBalanced(expr string) bool {
	s := stack.New[rune](1000)

	for _, r := range expr {
		if isOpener(r) {
			ok := s.Push(r)
			if !ok {
				panic("stack over capacity")
			}
		}
		if isCloser(r) {
			popped, ok := s.Pop()
			if !ok || popped != openerFor(r) {
				return false
			}
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
