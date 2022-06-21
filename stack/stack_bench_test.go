package stack_test

import (
	"go-katas/stack"
	"testing"
)

func BenchmarkName(b *testing.B) {
	s := stack.New[int](1_000_000)
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		for n := 0; n < 5; n++ {
			for j := 0; j < 500_000; j++ {
				s.Push(j)
			}
			for k := 0; k < 400_000; k++ {
				s.Pop()
			}
		}

	}
}
