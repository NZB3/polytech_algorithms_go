package useless

func InsertAt(a []int, index int, value int) []int {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func RemoveAt(a []int, index int) []int {
	return append(a[:index], a[index+1:]...)
}

func Append(a []int, value int) []int {
	return append(a, value)
}

func Prepend(a []int, value int) []int {
	return append([]int{value}, a...)
}
