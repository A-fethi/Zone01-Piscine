package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	result := ""
	for _, i := range os.Args[1:] {
		result = result + i + "\n"
	}
	myRune := []rune(result)
	for i := 0; i < len(myRune); i++ {
		z01.PrintRune(myRune[i])
	}
}
