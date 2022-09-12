package piscine

func squart(n int) int {
	result := 1
	for i := 1; i*i < n; i++ {
		result++
	}
	return result
}

func IsPrime(nb int) bool {
	var result bool = true
	if nb == 2 {
		return true
	}
	if nb > 0 && nb != 1 && nb <= 127 {
		for i := 2; i <= squart(nb); i++ {
			if nb%i == 0 {
				return false
			}
		}
	} else {
		result = false
	}
	return result
}
