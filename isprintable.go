package piscine

func IsPrintable(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] <= 13 {
			return false
		}
	}
	return true
}
