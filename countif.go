package piscine

func CountIf(f func(string) bool, tab []string) int {
	result := 0
	for i := 0; i < len(tab); i++ {
		if f(string(tab[i])) {
			result++
		}
	}
	return result
}
