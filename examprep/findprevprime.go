package main

import "fmt"

func isPrime(nb int) bool {
	if nb%2 == 0 || nb < 0 || nb == 1 {
		return false
	}
	return true
}

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	if isPrime(nb) {
		return (nb)
	}
	return FindPrevPrime(nb - 1)
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
