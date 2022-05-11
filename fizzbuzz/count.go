package fizzbuzz

import "strconv"

func Count(n int) []string {
	lines := []string{}
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			lines = append(lines, "FizzBuzz")
			continue
		}
		if i%5 == 0 {
			lines = append(lines, "Buzz")
			continue
		}
		if i%3 == 0 {
			lines = append(lines, "Fizz")
			continue
		}
		lines = append(lines, strconv.Itoa(i))
	}
	return lines
}
