package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	array := make([]int, max-min)
	for i := min; i < max; i++ {
		array[i-min] = i
	}
	return array
}
