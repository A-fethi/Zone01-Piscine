package main

import (
	"fmt"
	"strconv"
)

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	res := ""
	count := 1

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count++
		} else {
			res += strconv.Itoa(count) + string(s[i])
			count = 1
		}
	}
	res += strconv.Itoa(count) + string(s[len(s)-1])

	return res
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
