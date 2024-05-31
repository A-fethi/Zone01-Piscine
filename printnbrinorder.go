package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var result []int
	if n == 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		a := n % 10
		n /= 10
		result = append(result, a)
	}
	for i := 0; i <= 9; i++ {
		for j := range result {
			if result[j] == i {
				z01.PrintRune('0' + rune(result[j]))
			}
		}
	}
}
