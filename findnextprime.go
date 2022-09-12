package piscine

func isprime(nb int) bool {
	var result bool = true
	if nb == 2 {
		return true
	}
	if nb > 0 && nb != 1 {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
	} else {
		result = false
	}
	return result
}

func FindNextPrime(nb int) int {
	if isprime(nb) {
		return nb
	}
	return FindNextPrime(nb + 1)
}
