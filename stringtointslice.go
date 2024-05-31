package piscine

func StringToIntSlice(str string) []int {
	var slice []int
	for _, i := range str {
		slice = append(slice, int(i))
	}
	return slice
}
