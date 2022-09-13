package piscine

func NRune(s string, n int) rune {
	s_slice := []rune(s)
	if n > 0 && n <= len(s_slice) {
		return s_slice[n-1]
	} else {
		return 0
	}
}
