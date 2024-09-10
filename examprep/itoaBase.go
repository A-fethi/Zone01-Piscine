package main

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	digits := "0123456789ABCDEF"
	result := ""
	isNeg := value < 0
	if isNeg {
		value = -value
	}

	for value > 0 {
		rem := value % base
		result = string(digits[rem]) + result
		value = value / base
	}
	if result == "" {
		result = "0"
	}

	if isNeg {
		result = "-" + result
	}
	return result
}
