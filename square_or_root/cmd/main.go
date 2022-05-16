package main

import (
	"fmt"
	"math"
)

func main() {
	got := squareOrRoot([]int{4, 3, 9, 7, 2, 1})
	fmt.Println(got)
}

func squareOrRoot(list []int) []int {
	newlist := make([]int, len(list))
	for i, num := range list {
		sqrt := math.Sqrt(float64(num))
		if isWhole(sqrt) {
			newlist[i] = int(sqrt)
		} else {
			newlist[i] = num * num
		}
	}
	return newlist
}

func isWhole(fnum float64) bool {
	return float64(int(fnum)) == fnum
}
