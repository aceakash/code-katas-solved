package game_of_threes

func Count(num int) []int {
	nums := []int{}
	nums = append(nums, num)
	for num != 1 {
		if num%3 == 0 {
			num = num / 3
			nums = append(nums, num)
			continue
		}
		if num%3 == 1 {
			num -= 1
			nums = append(nums, num)
			num = num / 3
			nums = append(nums, num)
			continue
		}
		num += 1
		nums = append(nums, num)
		num = num / 3
		nums = append(nums, num)
	}
	return nums
}
