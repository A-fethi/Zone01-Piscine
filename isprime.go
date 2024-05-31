package piscine

func IsPrime(nb int) bool {
	if nb%2 == 0 || nb < 0 || nb == 1 {
		return false
	}
	return true
}
