package piscine

func ToUpper(s string) string {
	str := ""
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			str += string(i - 32)
		} else {
			str += string(i)
		}
	}
	return str
}
