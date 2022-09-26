package piscine

func Any(f func(string) bool, a []string) bool {
	b := []bool{}
	for i := 0; i < len(a); i++ {
		d := true
		for j := 0; j < len(a[i]); j++ {
			if !f(string(a[i][j])) {
				d = false
				break
			}
		}
		b = append(b, d)
		for i := 0; i < len(b); i++ {
			if b[i] {
				return true
			}
		}
	}
	return false
}
