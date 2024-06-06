package main

import "fmt"

func IsAlpha(str string) bool {
	for _, i := range str {
		if (i < 'a' || i > 'z') && (i < 'A' || i > 'Z') {
			return false
		}
	}
	return true
}

func IsUpper(str rune) bool {
	return str >= 'A' && str <= 'Z'
}

func CamelToSnakeCase(s string) string {
	result := ""
	if s == "" || !IsAlpha(s) {
		return s
	}
	for i := 0; i < len(s); i++ {
		if IsUpper(rune(s[i])) || (i == 0 && !IsUpper(rune(s[i]))) {
			result = string(s[i]) + result
		}
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase(""))
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
