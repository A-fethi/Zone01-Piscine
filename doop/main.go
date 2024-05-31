package main

import (
	"os"

	"github.com/01-edu/z01"
)

func strRev(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func atoi(str string) int {
	sign := 1
	if str == "" {
		return 0
	}
	if str[0] == '-' || str[0] == '+' {
		if str[0] == '-' {
			sign *= -1
		}
		str = str[1:]
	}
	str = strRev(str)
	div := 1
	result := 0
	for _, i := range str {
		nb := int(i - 48)
		if nb > 9 || nb < 0 {
			return 0
		}
		result += nb * div
		div *= 10
	}
	return result * sign
}

func printStr(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) == 4 {
		num1 := atoi(os.Args[1])
		num2 := atoi(os.Args[3])
		if os.Args[2] == "+" {
			myRune := '0' + rune(num1+num2)
			z01.PrintRune(myRune)
			z01.PrintRune('\n')
		}
		if os.Args[2] == "-" {
			myRune := '0' + rune(num1+num2)
			z01.PrintRune(myRune)
			z01.PrintRune('\n')
		}
		if os.Args[2] == "*" {
			myRune := '0' + rune(num1+num2)
			z01.PrintRune(myRune)
			z01.PrintRune('\n')
		}
		if os.Args[2] == "/" {
			if os.Args[3] != "0" {
				myRune := '0' + rune(num1+num2)
				z01.PrintRune(myRune)
				z01.PrintRune('\n')
			} else {
				printStr("No division by 0")
			}
		}
		if os.Args[2] == "%" {
			if os.Args[3] == "0" {
				printStr("No modulo by 0")
			} else {
				myRune := '0' + rune(num1+num2)
				z01.PrintRune(myRune)
				z01.PrintRune('\n')
			}
		}
	}
}
