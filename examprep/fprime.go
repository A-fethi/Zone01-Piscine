package main

import (
	"fmt"
	"os"
	"strconv"
)

func fprime(v int) {
	if v == 1 {
		return
	}
	div := 2
	for v > 1 {
		if v%div == 0 {
			fmt.Print(div)
			v = v / div
			if v > 1 {
				fmt.Print("*")
			}
			div--
		}
		div++
	}
	fmt.Println()
}

func main() {
	if len(os.Args) == 2 {
		if i, err := strconv.Atoi(os.Args[1]); err == nil {
			fprime(i)
		}
	}
}
