package main

import (
	"os"

	"github.com/01-edu/z01"
)

func IsIn(s rune, arr []rune) bool {
	for _, i := range arr {
		if s == i {
			return true
		}
	}
	return false
}

func main() {
	var result []rune
	arg1 := []rune(os.Args[1])
	arg2 := []rune(os.Args[2])
	if len(os.Args) == 3 {
		for _, i := range arg1 {
			if !IsIn(i, result) {
				result = append(result, i)
			}
		}
		for _, i := range arg2 {
			if !IsIn(i, result) {
				result = append(result, i)
			}
		}
		for _, i := range result {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
