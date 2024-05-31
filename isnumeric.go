package piscine

func IsNumeric(s string) bool {
	myRune := []rune(s)
	for i := 0; i < len(myRune); i++ {
		if myRune[i] < '0' || myRune[i] > '9' {
			return false
		}
	}
	return true
}
