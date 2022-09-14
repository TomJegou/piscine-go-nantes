package piscine

func isLowerForCapitalize(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] < 'a' || s_S[i] > 'z' {
			return false
		}
	}
	return true
}

func isUpperForCapitalize(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] < 'A' || s_S[i] > 'Z' {
			return false
		}
	}
	return true
}

func isLetter(s string) bool {
	if len(s) == 0 {
		return true
	}
	result := false
	for i := 0; i < len(s); i++ {
		if isLowerForCapitalize(string(s[i])) || isUpperForCapitalize(string(s[i])) {
			result = true
		} else {
			return false
		}

	}
	return result
}

func Capitalize(s string) string {
	slice_string := []byte(s)
	search_mode := true
	for i := 0; i < len(s); i++ {
		if search_mode && isLetter(string(slice_string[i])) {
			if isLowerForCapitalize(string(slice_string[i])) {
				search_mode = false
				slice_string[i] -= 32
			} else {
				search_mode = false
			}
		} else if !isLetter(string(slice_string[i])) {
			search_mode = true
		}
	}
	return string(slice_string)
}
