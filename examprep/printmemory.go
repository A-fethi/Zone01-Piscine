package main

import "github.com/01-edu/z01"

func convert(arr [10]byte) {
	base := "0123456789abcdef"
	i := 0
	for ; i < len(arr); i++ {
		n := int(arr[i]) / len(base)
		z01.PrintRune(rune(base[n]))
		index := int(arr[i]) % len(base)
		z01.PrintRune(rune(base[index]))
		if !(i == 3 || i == 7) {
			z01.PrintRune(' ')
		}
		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')
}

func PrintMemory(arr [10]byte) {
	convert(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			z01.PrintRune(rune(arr[i]))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
