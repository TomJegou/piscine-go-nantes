package piscine

func AlphaCount(s string) int {
	result := 0
	s_byte := 0
	s_s := []byte(s)
	for i := 0; i < len(s); i++ {
		s_byte = int(s_s[i])
		if s_byte >= 65 && s_byte <= 90 {
			result++
		} else if s_byte >= 97 && s_byte <= 122 {
			result++
		}
	}
	return result
}
