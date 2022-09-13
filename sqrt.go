package piscine

func Sqrt(nb int) int {
	var result int
	if nb > 0 {
		count := 1
		for i := 1; i*i <= nb; i++ {
			count++
		}
		nb -= count * count
		if nb == 0 {
			result = count
		}
	} else {
		result = 0
	}
	return result
}
