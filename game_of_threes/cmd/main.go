package main

import (
	"fmt"
	"go-katas/game_of_threes"
	"math"
)

func main() {
	res := game_of_threes.Count(math.MaxInt64)
	for _, num := range res {
		fmt.Println(num)
	}
}
