package piscine

func islower(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] < 'a' || s_S[i] > 'z' {
			return false
		}
	}
	return true
}

func ToUpper(s string) string {
	s_byte := []byte(s)
	for i := 0; i < len(s); i++ {
		if islower(string(s[i])) {
			s_byte[i] -= 32
		}
	}
	return string(s_byte)
}
