package piscine

func isupper(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] < 'A' || s_S[i] > 'Z' {
			return false
		}
	}
	return true
}

func ToLower(s string) string {
	s_byte := []byte(s)
	for i := 0; i < len(s); i++ {
		if isupper(string(s[i])) {
			s_byte[i] += 32
		}
	}
	return string(s_byte)
}
