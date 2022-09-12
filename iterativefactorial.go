package piscine

func IterativeFactorial(nb int) int {
	var result int = 1
	if nb >= -128 && nb <= 127 {
		for i := 1; i <= nb; i++ {
			result *= i
		}
	} else {
		result = 0
	}
	return result
}
