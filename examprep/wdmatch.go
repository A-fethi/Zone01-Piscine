package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	arg1 := os.Args[1]
	arg2 := os.Args[2]

	j := 0

	for i := 0; i < len(arg1); i++ {
		for j < len(arg2) && arg1[i] != arg2[j] {
			j++
		}

		if j == len(arg2) {
			return
		}

		j++
	}

	for _, r := range arg1 {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
