package main

import (
	"fmt"
)

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}

func FifthAndSkip(str string) string {
	s := ""
	count := 0
	if len(str) < 5 {
		return "Invalid Input\n"
	}
	for _, c := range str {
		if count != 0 && count%5 == 0 {
			s += " "
			count = 0
		} else if c != ' ' {
			s += string(c)
			count++
		}
	}
	return s + "\n"
}

// func FifthAndSkip(str string) string {
// 	// Handle empty input
// 	if str == "" {
// 		return "\n"
// 	}

// 	// Handle invalid input (less than 5 characters)
// 	if len(str) < 5 {
// 		return "Invalid Input\n"
// 	}

// 	// Initialize a variable to hold the result
// 	var result string
// 	count := 0

// 	// Iterate over the string, ignoring spaces
// 	for i := 0; i < len(str); i++ {
// 		// Skip spaces
// 		if str[i] == ' ' {
// 			continue
// 		}

// 		// Add the character to the result
// 		result += string(str[i])
// 		count++

// 		// Every 5 characters, add a space
// 		if count == 5 {
// 			result += " "
// 			count = 0
// 		}
// 	}

// 	// Remove the trailing space if exists
// 	if len(result) > 0 && result[len(result)-1] == ' ' {
// 		result = result[:len(result)-1]
// 	}

// 	return result
// }
