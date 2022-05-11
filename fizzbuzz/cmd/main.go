package main

import (
	"fmt"
	"go-katas/fizzbuzz"
)

func main() {
	lines := fizzbuzz.Count(100)
	for _, line := range lines {
		fmt.Println(line)
	}
}
