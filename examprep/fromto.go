package piscine

import "strconv"

func FromTo(from, to int) string {
	result := ""

	if from > 99 || from < 0 || to > 99 || to < 0 {
		return "Invalid\n"
	} else if from == to && from < 10 {
		return "0" + strconv.Itoa(from) + "\n"
	}
	if from > to {
		for i := from; i >= to; i-- {
			if i < 10 {
				result += "0"
			}
			result += strconv.Itoa(i)
			if i-1 >= to {
				result += ", "
			}
		}
		return result + "\n"
	}
	for i := from; i <= to; i++ {
		if i < 10 {
			result += "0"
		}
		result += strconv.Itoa(i)
		if i+1 <= to {
			result += ", "
		}
	}
	return result + "\n"
}