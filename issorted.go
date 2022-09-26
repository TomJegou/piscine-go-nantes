package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	l := []int{}
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) != 0 {
			l = append(l, f(a[i-1], a[i]))
		}
	}
	for i := 1; i < len(l); i++ {
		if l[i] != l[0] {
			return false
		}
	}
	return true
}
