package piscine

func found_biggest(t []byte) int {
	index_biggest := 1
	for i := 0; i < len(t); i++ {
		if t[i] > t[index_biggest] {
			index_biggest = i
		}
	}
	return index_biggest
}

func substarct(t []byte, restult byte) byte {
	if len(t) == 1 {
		return t[0]
	}
	i := found_biggest(t)
	a := t[i]
	t = append(t[:i], t[i+1:]...)
	return restult - substarct(t, a)
}

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if substarct([]byte(a), 0) < substarct([]byte(b), 0) {
		return -1
	} else {
		return 1
	}
}
