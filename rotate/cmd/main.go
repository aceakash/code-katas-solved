package main

import (
	"fmt"
	"go-katas/rotate"
)

func main() {
	list := []int{7, 8, 9}
	rotate.RotateInPlace(list, 13)
	fmt.Println(list)
}
