package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	myRune := []rune(os.Args[0])
	for i := 2; i < len(myRune); i++ {
		z01.PrintRune(myRune[i])
	}
	z01.PrintRune('\n')
}
