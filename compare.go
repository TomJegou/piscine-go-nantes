package piscine

func Compare(a, b string) int {
	a_s := []byte(a)
	b_s := []byte(b)
	len_min := 0
	if a == b {
		return 0
	} else {
		if len(a) < len(b) {
			len_min = len(a)
		} else {
			len_min = len(b)
		}
		for i := 0; i < len_min; i++ {
			if int(a_s[i]) < int(b_s[i]) {
				return -1
			}
		}
	}
	return 1
}
