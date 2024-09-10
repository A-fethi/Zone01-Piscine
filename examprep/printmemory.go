package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	base := "0123456789abcdef"

	for i := 0; i < len(arr); i++ {
		div := string(base[arr[i]/16])
		mod := string(base[arr[i]%16])
		z01.PrintRune([]rune(div)[0])
		z01.PrintRune([]rune(mod)[0])
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else if i != len(arr)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
	for _, char := range arr {
		if char > ' ' && char < '~' {
			z01.PrintRune(rune(char))
			continue
		}
		z01.PrintRune('.')
	}
	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
