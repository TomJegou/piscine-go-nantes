package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	d := []int{}
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) != 0 {
			d = append(d, f(a[i-1], a[i]))
		}
	}
	if len(d) > 0 {
		sign := d[0] > 0
		for i := 1; i < len(d); i++ {
			if sign {
				if d[i] < 0 {
					return false
				}
			} else if !sign {
				if d[i] > 0 {
					return false
				}
			}
		}
	}
	return true
}
