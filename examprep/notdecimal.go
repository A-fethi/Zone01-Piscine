package main

import "fmt"

func NotDecimal(s string) string {
	var result string

	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			return s
		} else if char >= 'a' && char <= 'z' {
			return s
		}
	}

	dec := atoi(s)
	result = itoa(dec)
	return result

}

func atoi(s string) int {
	res := 0
	sign := 1

	for i, char := range s {
		if i == 0 && char == '-' {
			sign *= -1
		} else if char >= '0' && char <= '9' {
			res = res*10 + int(char-'0')
		} else if char == '.' {
			continue
		}
	}
	return res * sign
}

func itoa(num int) string {
	res := ""
	sign := ""

	if num < 0 {
		num *= -1
		sign = "-"
	}

	for num > 0 {
		digit := num % 10
		res = string(digit+'0') + res
		num /= 10
	}
	return sign + res
}

func main() {
	fmt.Println(NotDecimal("0.1"))
	fmt.Println(NotDecimal("174.2"))
	fmt.Println(NotDecimal("0.1255"))
	fmt.Println(NotDecimal("1.20525856"))
	fmt.Println(NotDecimal("-0.0f00d00"))
	fmt.Println(NotDecimal(""))
	fmt.Println(NotDecimal("51.3+95.9"))
	fmt.Println(NotDecimal("1952"))
	fmt.Println(NotDecimal("165"))
	fmt.Println(NotDecimal("-.19-52"))
	fmt.Println(NotDecimal("415.458"))
	fmt.Println(NotDecimal("56s44"))
	fmt.Println(NotDecimal("00.02"))
}