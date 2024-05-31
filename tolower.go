package piscine

func ToLower(s string) string {
	str := ""
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			str += string(i + 32)
		} else {
			str += string(i)
		}
	}
	return str
}
