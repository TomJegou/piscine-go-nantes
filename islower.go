package piscine

func IsLower(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] < 'a' || s_S[i] > 'z' {
			return false
		}
	}
	return true
}
