package piscine

func RecursiveFactorial(nb int) int {
	if nb <= 0 || nb > 20 {
		if nb == 0 {
			return 1
		} else {
			return 0
		}
	}
	var result int
	result = nb * RecursiveFactorial(nb-1)
	return result
}
