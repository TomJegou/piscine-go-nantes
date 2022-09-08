package piscine

// func pow(m int, n int) int {
// 	if n == 0 {
// 		return 1
// 	}

// 	result := m
// 	for i := 0; i < n-1; i++ {
// 		result = result * m
// 	}
// 	return result
// }

func BasicAtoi2(s string) int {
	var result int
	aStringChangeable := []byte(s)
	count_pow := 0
	for i := 0; i < len(s); i++ {
		ascii_digit := int(aStringChangeable[len(s)-count_pow-1])
		if ascii_digit >= 48 && ascii_digit <= 57 {
			result = result + (ascii_digit-48)*pow(10, i)
			count_pow++
		} else {
			return 0
		}
	}
	return result
}
