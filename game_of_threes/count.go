package game_of_threes

func Count(num int) []int {
	nums := []int{num}
	for num != 1 {
		if num%3 == 2 {
			num += 1
			nums = append(nums, num)
		}
		if num%3 == 1 {
			num -= 1
			nums = append(nums, num)
		}
		num = num / 3
		nums = append(nums, num)
		continue
	}
	return nums
}
