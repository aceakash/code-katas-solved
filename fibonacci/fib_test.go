package main_test

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

func Test_Can_Compute_The_Nth_Fibonacci_Number(t *testing.T) {
	//t.Run("works for small numbers", func(t *testing.T) {
	//	assert.Equal(t, 0, Fib(0))
	//	assert.Equal(t, 3, Fib(4))
	//	assert.Equal(t, 610, Fib(15))
	//})

	t.Run("works for bigger numbers", func(t *testing.T) {
		knownFibs := map[int]int{}
		assert.Equal(t, 1134903170, Fib(45, knownFibs))
		assert.Equal(t, 1548008755920, Fib(60, knownFibs))
		assert.Equal(t, 7540113804746346429, Fib(92, knownFibs))
	})
}

func Fib(n int, knownFibs map[int]int) int {
	if f, ok := knownFibs[n]; ok {
		return f
	}
	if n <= 0 {
		knownFibs[n] = 0
		return 0
	}
	if n == 1 {
		knownFibs[n] = 1
		return 1
	}
	knownFibs[n] = Fib(n-1, knownFibs) + Fib(n-2, knownFibs)
	return knownFibs[n]
}
