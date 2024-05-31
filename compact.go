package piscine

func Compact(ptr *[]string) int {
	length := 0
	for _, i := range *ptr {
		if i != "" {
			length++
		}
	}
	j := 0
	arr := make([]string, length)
	for _, i := range *ptr {
		if i != "" {
			arr[j] = i
			j++
		}
	}
	*ptr = arr
	return length
}
