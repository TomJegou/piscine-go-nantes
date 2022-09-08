package piscine

func StrRev(s string) string {
	aStringChangeable := []byte(s)
	for i := 0; i < len(s); i++ {
		aStringChangeable[i] = s[len(s)-i-1]
	}
	result := string(aStringChangeable)
	return result
}
