package main

import (
	"fmt"
	"strings"
)

func WordFlip(str string) string {
	// Trim leading and trailing spaces
	str = strings.TrimSpace(str)

	// If the string is empty after trimming, return "Invalid Output"
	if str == "" {
		return "Invalid Output\n"
	}

	// Split the string into words based on spaces, automatically ignoring multiple spaces
	words := strings.Fields(str)

	// Reverse the order of words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Join the words back into a single string, separated by a space
	result := strings.Join(words, " ")

	// Return the result followed by a newline
	return result + "\n"
}

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
