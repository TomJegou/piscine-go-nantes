package piscine

func Sqrt(nb int) int {
	if nb > 0 {
		count := 1
		for i := 1; i*i < nb; i++ {
			count++
		}
		if count*count == nb {
			return count
		}
	}
	return 0
}
