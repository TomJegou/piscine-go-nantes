package piscine

func IterativeFactorial(nb int) int {
	var result int = 1
	for i := 0; i < nb; i++ {
		result *= i
	}
	return result
}
