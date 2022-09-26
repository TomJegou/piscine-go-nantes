package piscine

func Any(f func(string) bool, a []string) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if !f(string(a[i][j])) {
				return false
			}
		}
	}
	return true
}
