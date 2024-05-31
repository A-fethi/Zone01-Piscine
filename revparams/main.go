package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	result := ""
	for i := len(os.Args) - 1; i > 0; i-- {
		result = result + os.Args[i] + "\n"
	}
	myRune := []rune(result)
	for i := 0; i < len(myRune); i++ {
		z01.PrintRune(myRune[i])
	}
}
