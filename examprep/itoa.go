package main

import "fmt"

func Itoa(num int) string {
	result := ""
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num *= -1
		for num > 0 {
			result = string(rune(num%10+48)) + result
			num /= 10
		}
		result = "-" + result
	}
	for num > 0 {
		result = string(rune(num%10+48)) + result
		num /= 10
	}
	return result
}

func main() {
    fmt.Println(Itoa(+12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}
