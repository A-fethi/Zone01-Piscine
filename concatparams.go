package piscine

func ConcatParams(args []string) string {
	result := ""
	for i := 0; i < len(args); i++ {
		if i < len(args)-1 {
			result += args[i] + "\n"
		} else {
			result += args[i]
		}
	}
	return result
}
