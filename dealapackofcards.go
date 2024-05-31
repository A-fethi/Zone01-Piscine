package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		start := i * 3
		fmt.Printf("Player %d: %d, %d, %d", i+1, deck[start], deck[start+1], deck[start+2])
		z01.PrintRune('\n')
	}
}
