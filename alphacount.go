package piscine

func AlphaCount(s string) int {
	myRune := []rune(s)
	count := 0
	for i := 0; i < len(myRune); i++ {
		if myRune[i] >= 'a' && myRune[i] <= 'z' || myRune[i] >= 'A' && myRune[i] <= 'Z' {
			count++
		}
	}
	return count
}
