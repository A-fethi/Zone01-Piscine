package piscine

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] != '\n' && s[i] != ' ' && s[i] != '\t' {
			word += string(s[i])
		} else {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
