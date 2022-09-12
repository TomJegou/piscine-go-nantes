package piscine

func IterativePower(nb int, power int) int {
	var result int
	if power >= 0 {
		result = nb
		for i := 0; i < power-1; i++ {
			result *= nb
		}
	}
	return result
}
