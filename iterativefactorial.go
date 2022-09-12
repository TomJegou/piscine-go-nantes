package piscine

func IterativeFactorial(nb int) int {
	var result int = 1
	if nb > 0 {
		for i := 0; i < nb; i++ {
			result *= i
		}
	} else {
		result = 0
	}

	return result
}
