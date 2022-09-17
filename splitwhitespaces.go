package piscine

func _split_custom(s string, t string) []string {
	string_temp := ""
	result := []string{}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == t {
			if len(string_temp) > 0 {
				result = append(result, string_temp)
				string_temp = ""
			}
			s = s[:i] + s[i+1:]
			i -= 1
		} else {
			string_temp += string(s[i])
		}
	}
	if len(string_temp) > 0 {
		result = append(result, string_temp)
	}
	return result
}

func SplitWhiteSpaces(s string) []string {
	return _split_custom(s, " ")
}
