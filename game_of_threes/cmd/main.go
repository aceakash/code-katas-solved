package main

import "fmt"

func main() {
	nums := count(67)
	for _, num := range nums {
		fmt.Println(num)
	}
}

func count(num int) []int {
	nums := []int{}
	for num != 1 {
		nums = append(nums, num)
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
	nums = append(nums, num)
	return nums
}
