package piscine

func lastrune(s string) rune {
	s_slice := []rune(s)
	return s_slice[len(s)-1]
}
