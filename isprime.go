package piscine

func squart(n int) int {
	result := 1
	for i := 1; i*i < n; i++ {
		result++
	}
	return result
}

func IsPrime(nb int) bool {
	s := squart(nb)
	var result bool = true
	if nb > 0 && nb != 1 {
		for i := 2; i <= s; i++ {
			if i*i == nb {
				result = false
			} else {
				continue
			}
		}
	} else {
		result = false
	}

	return result
}
