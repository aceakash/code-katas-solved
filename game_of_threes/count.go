package game_of_threes

func Count(num int) []int {
	if num == 1 {
		return []int{1}
	}
	next := nextNumber(num)
	return append([]int{num}, Count(next)...)
}

func nextNumber(num int) (next int) {
	if num%3 == 2 {
		return num + 1
	}
	if num%3 == 1 {
		return num - 1
	}
	return num / 3
}
