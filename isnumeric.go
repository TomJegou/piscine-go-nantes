package piscine

func IsNumeric(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] < '0' || s_S[i] > '9' {
			return false
		}
	}
	return true
}
