package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	l := []int{}
	for i := 1; i < len(a); i++ {
		l = append(l, f(a[i-1], a[i]))
	}
	b := l[0]
	for i := 1; i < len(l); i++ {
		if l[i] != b {
			return false
		}
	}
	return true
}
