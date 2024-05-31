package piscine

func NRune(s string, n int) rune {
	myRune := []rune(s)
	for i := 0; i < len(myRune); i++ {
		if i == n-1 {
			return myRune[i]
		}
	}
	return 0
}
