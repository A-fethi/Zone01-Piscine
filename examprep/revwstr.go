package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if there is exactly one argument
	if len(os.Args) != 2 {
		return
	}

	// Get the input string (the second argument, since os.Args[0] is the program name)
	input := os.Args[1]

	// Split the input string into words
	words := strings.Split(input, " ")

	// Reverse the words slice
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Print the reversed words as a single string
	fmt.Println(strings.Join(words, " "))
}
