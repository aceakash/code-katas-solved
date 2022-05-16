package rotate

func Rotate(list []int, by int) []int {
	if by < 0 {
		panic("by cannot be negative")
	}
	if len(list) <= 1 || by == 0 {
		return list
	}

	newList := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[(i+by)%len(list)]
	}
	return newList
}

func RotateInPlace(list []int, by int) {
	if by < 0 {
		panic("by cannot be negative")
	}
	if len(list) <= 1 || by == 0 {
		return
	}

	for i := 0; i < by; i++ {
		rotateByOne(list)
	}
}

func rotateByOne(list []int) {
	temp := list[0]
	for i := 0; i < len(list)-1; i++ {
		list[i] = list[i+1]
	}
	list[len(list)-1] = temp
}
