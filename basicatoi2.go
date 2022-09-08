package piscine

func BasicAtoi2(s string) int {
	var result int
	aStringChangeable := []byte(s)
	for i := 0; i < len(s); i++ {
		ascii_digit := int(aStringChangeable[len(s)-i-1])
		if ascii_digit >= 48 && ascii_digit <= 57 {
			result = result + (ascii_digit-48)*pow(10, i)
		} else {
			return 0
		}
	}
	return result
}
