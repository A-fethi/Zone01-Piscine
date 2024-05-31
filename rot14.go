package piscine

func Rot14(s string) string {
	result := ""
	for _, i := range s {
		if i <= 'l' && i >= 'a' || i <= 'L' && i >= 'A' {
			result += string(i + 14)
		} else if i > 'l' && i <= 'z' || i > 'L' && i <= 'Z' {
			result += string(i - 12)
		} else {
			result += string(i)
		}
	}
	return result
}
