package piscine

func JumpOver(str string) string {
	result := ""
	myRune := []rune(str)
	if str == "" {
		return "\n"
	}
	for i := 2; i < len(str); i += 3 {
		result += string(myRune[i])
	}
	return result + "\n"
}
