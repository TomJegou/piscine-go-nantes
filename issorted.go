package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	d := f(a[0], a[1])
	for i := 2; i < len(a); i++ {
		if f(a[i-1], a[i]) != d && f(a[i-1], a[i]) != 0 {
			return false
		}
	}
	return true
}
