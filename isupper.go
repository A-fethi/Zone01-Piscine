package piscine

func IsUpper(s string) bool {
	myRune := []rune(s)
	for i := 0; i < len(myRune); i++ {
		if myRune[i] <= 'A' || myRune[i] >= 'Z' {
			return false
		}
	}
	return true
}
