package piscine

func isLower(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] < 'a' || s_S[i] > 'z' {
			return false
		}
	}
	return true
}

func isUpper(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] < 'A' || s_S[i] > 'Z' {
			return false
		}
	}
	return true
}

func isNum(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] < '0' || s_S[i] > '9' {
			return false
		}
	}
	return true
}

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return true
	}
	result := false
	for i := 0; i < len(s); i++ {
		if isLower(string(s[i])) || isUpper(string(s[i])) || isNum(string(s[i])) {
			result = true
		} else {
			return false
		}
	}
	return result
}
