package main

//import "fmt"

func strRev(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	sign := 1
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign *= -1
		}
		s = s[1:]
	}
	result := 0
	div := 1
	s = strRev(s)
	for _, i := range s {
		nb := int(i - 48)
		if nb > 9 || nb < 0 {
			return 0
		}
		result += nb * div
		div *= 10
	}
	return result * sign
}

/*func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}*/
