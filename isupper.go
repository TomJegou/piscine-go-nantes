package piscine

func IsUpper(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] < 'A' || s_S[i] > 'Z' {
			return false
		}
	}
	return true
}
