package main

import "fmt"

func FibRecursive(n int, knownFibs map[int]int) int {
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
	nextFib := FibRecursive(n-1, knownFibs) + FibRecursive(n-2, knownFibs)

	if nextFib < 0 {
		msg := fmt.Sprintf("integer overflow when calculating next fibonacci sequence number #%d. Prev numbers were %d (#%d) and %d (#%d)", n, FibRecursive(n-1, knownFibs), n-1, FibRecursive(n-2, knownFibs), n-2)
		panic(msg)
	}
	knownFibs[n] = nextFib
	return knownFibs[n]
}

func FibIterative(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a := 0
	b := 1
	for i := 2; i <= n; i++ {
		c := a + b
		if c < 0 {
			msg := fmt.Sprintf("integer overflow when calculating next fibonacci sequence number #%d. Prev numbers were %d (#%d) and %d (#%d)", i, a, i-2, b, i-1)
			panic(msg)
		}
		a, b = b, c
	}

	return b

}
