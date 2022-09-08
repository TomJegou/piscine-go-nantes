package piscine

func Atoi(s string) int {
	var result int
	aStringChangeable := []byte(s)
	count_pow := 0
	if int(aStringChangeable[0]) == 43 || int(aStringChangeable[0]) == 45 {
		if int(aStringChangeable[1]) >= 48 && int(aStringChangeable[1]) <= 57 {
			for i := 1; i < len(s); i++ {
				ascii_digit := int(aStringChangeable[len(s)-count_pow-1])
				if ascii_digit >= 48 && ascii_digit <= 57 {
					result = result + (ascii_digit-48)*pow(10, i)
					count_pow++
				} else {
					return 0
				}
			}
		} else {
			return 0
		}
	} else {
		for i := 0; i < len(s); i++ {
			ascii_digit := int(aStringChangeable[len(s)-count_pow-1])
			if ascii_digit >= 48 && ascii_digit <= 57 {
				result = result + (ascii_digit-48)*pow(10, i)
				count_pow++
			} else {
				return 0
			}
		}
	}
	if int(aStringChangeable[0]) == 45 {
		result *= -1
	}
	return result
}
