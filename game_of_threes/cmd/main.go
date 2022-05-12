package main

import (
	"fmt"
	"go-katas/game_of_threes"
)

func main() {
	nums := game_of_threes.Count(76543)
	for _, num := range nums {
		fmt.Println(num)
	}
}
