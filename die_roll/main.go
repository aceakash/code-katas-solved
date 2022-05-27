package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rollFreq := RollDie(6, 1_000_000, r)
	fmt.Println(rollFreq)
}

func RollDie(faceCount int, rollCount int, r *rand.Rand) map[int]int {
	freq := make(map[int]int, faceCount)
	for i := 0; i < rollCount; i++ {
		face := r.Intn(faceCount) + 1
		freq[face] += 1
	}
	return freq
}
