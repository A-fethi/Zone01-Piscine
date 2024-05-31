package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	result := ""
	var burger food
	var chips food
	var nuggets food
	burger.preptime = 15
	chips.preptime = 10
	nuggets.preptime = 12
	for _, i := range order {
		result += string(i)
	}
	if result == "burger" {
		return burger.preptime
	}
	if result == "chips" {
		return chips.preptime
	}
	if result == "nuggets" {
		return nuggets.preptime
	}
	return 404
}
