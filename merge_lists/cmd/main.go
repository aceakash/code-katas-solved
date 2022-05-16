package main

import "fmt"

func main() {
	fmt.Println(merge([]int{1, 4, 6}, []int{2, 3, 5}))
}

func merge(a []int, b []int) []int {
	c := make([]int, len(a)+len(b))

	i, j, k := 0, 0, 0
	for k < len(a)+len(b) {
		if j >= len(b) || a[i] < b[j] {
			c[k] = a[i]
			i++
			k++
			continue
		}
		if i >= len(a) || b[j] < a[i] {
			c[k] = b[j]
			j++
			k++
			continue
		}
		c[k] = a[i]
		c[k+1] = b[j]
		i++
		j++
		k += 2
	}
	return c
}
