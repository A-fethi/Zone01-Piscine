package piscine

func IsLower(s string) bool {
	myRune := []rune(s)
	for i := 0; i < len(myRune); i++ {
		if myRune[i] < 'a' || myRune[i] > 'z' {
			return false
		}
	}
	return true
}
