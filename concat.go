package piscine

func Concat(str1 string, str2 string) string {
	s1 := []byte(str1)
	s2 := []byte(str2)
	s1 = append(s1, s2...)
	return string(s1)
}
