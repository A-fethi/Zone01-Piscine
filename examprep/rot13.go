package main

import "os"

func main() {
	result := ""
	if len(os.Args) == 2 {
		for _, i := range os.Args[1] {
			if i >= 'a' && i <= 'm' || i >= 'A' && i <= 'M' {
				result += string(i + 13)
			} else if i > 'm' && i <= 'z' || i > 'M' && i <= 'Z' {
				result += string(i - 13)
			} else {
				result += string(i)
			}
		}
		print(result)
		print("\n")
	}
}
