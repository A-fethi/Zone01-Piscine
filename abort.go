package piscine

func Abort(a, b, c, d, e int) int {
	arg := []int{a, b, c, d, e}
	for i := 0; i < len(arg); i++ {
		for j := 0; j < len(arg)-i-1; j++ {
			if arg[j] > arg[j+1] {
				arg[j], arg[j+1] = arg[j+1], arg[j]
			}
		}
	}
	return arg[2]
}
