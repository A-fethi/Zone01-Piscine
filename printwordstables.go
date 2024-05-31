package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(s []string) {
	for _, i := range s {
		for _, j := range i {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
