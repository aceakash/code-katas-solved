package main

import "fmt"

func main() {
	fmt.Println(LetterSum("43FG%$.<> microspectrophotometries"))
}

func LetterSum(s string) int {
	sum := 0
	for _, r := range s {
		if r < 'a' || r > 'z' {
			continue
		}
		sum += int(r) - 96
	}
	return sum
}
