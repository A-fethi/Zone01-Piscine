package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printNbr(n int) {
	if n >= 10 {
		nbr := n / 10
		mod := n % 10
		printNbr(nbr)
		printNbr(mod)
	} else {
		for i, j := 0, '0'; i <= n; i, j = i+1, j+1 {
			if i == n {
				z01.PrintRune(j)
			}
		}
	}
}

func main() {
	points := &point{}
	setPoint(points)

	Xresult := "x = "
	for _, i := range Xresult {
		z01.PrintRune(i)
	}
	printNbr(points.x)
	Yresult := ", y = "
	for _, i := range Yresult {
		z01.PrintRune(i)
	}
	printNbr(points.y)
	z01.PrintRune('\n')
}
