package piscine

func convert_int_to_string_for_atoi(start int, end int, str string) int {
	aStringChangeable := []byte(str)
	var result int
	for i := start; i >= end; i-- {
		ascii_digit := int(aStringChangeable[i])
		if ascii_digit >= 48 && ascii_digit <= 57 {
			result += (ascii_digit - 48) * pow(10, len(str)-1-i)
		} else {
			return 0
		}
	}
	return result
}

func Atoi(s string) int {
	var result int
	if len(s) > 0 {
		aStringChangeable := []byte(s)
		if int(aStringChangeable[0]) == 43 || int(aStringChangeable[0]) == 45 {
			if int(aStringChangeable[1]) >= 48 && int(aStringChangeable[1]) <= 57 {
				result = convert_int_to_string_for_atoi(len(s)-1, 1, s)
			} else {
				return 0
			}
		} else {
			result = convert_int_to_string_for_atoi(len(s)-1, 0, s)
		}
		if int(aStringChangeable[0]) == 45 {
			result *= -1
		}
	}
	return result
}
