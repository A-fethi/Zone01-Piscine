package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) == 1 {
		return true
	}
	if f(a[0], a[1]) < 0 {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		}
	} else {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		}
	}
	return true
}

/*func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{929173, 680386, 645379, 257272, 179965, -40947, -134423, -494271}
	a4 := []int{1}
	a5 := []int{0, 1, 2, 3, 4, 5}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)
	result3 := IsSorted(f, a3)
	result4 := IsSorted(f, a4)
	result5 := IsSorted(f, a5)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
}*/
