package main

import "fmt"

func main() {
	count(67)
}

func count(num int) {
	for num != 1 {
		fmt.Println(num)
		if num%3 == 0 {
			num = num / 3
			continue
		}
		if num%3 == 1 {
			num -= 1
			num = num / 3
			continue
		}
		num += 1
		num = num / 3
	}
	fmt.Println(num)
}
