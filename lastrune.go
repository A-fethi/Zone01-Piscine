package piscine

func LastRune(s string) rune {
	myRune := []rune(s)
	return myRune[len(s)-1]
}
