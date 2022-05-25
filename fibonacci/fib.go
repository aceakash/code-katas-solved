package main

import (
	"fmt"
	"math/big"
)

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

func FibBig(n int) string {
	if n == 0 {
		return "0"
	}
	if n == 1 {
		return "1"
	}
	a := big.NewInt(0)
	b := big.NewInt(1)

	// Initialize limit as 10^100
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(100), nil)

	for i := 2; i <= n; i++ {
		a.Add(a, b)
		if a.Cmp(&limit) > 0 {
			msg := fmt.Sprintf("encountered fib sequence number %s (#%d) which is over the configured limit of %s", a.String(), i, limit.String())
			panic(msg)
		}
		a, b = b, a
	}
	return b.String()
}
