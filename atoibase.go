package piscine

func powforAtoibase(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if nb == 1 || nb == 0 {
		return nb
	} else if power == 0 {
		return 1
	}
	a := powforAtoibase(nb, power/2)
	if power%2 == 0 {
		return a * a
	} else {
		return a * a * nb
	}
}

func switch_str_for_atoibase(a int, b int, t []byte) {
	c := t[a]
	t[a] = t[b]
	t[b] = c
}

func reverse_slice_byte_for_atoibase(t []byte) {
	j := len(t) - 1
	for i := 0; i < j; i++ {
		switch_str_for_atoibase(i, j, t)
		j--
	}
}

func found_index(s byte, base string) int {
	index := 0
	for i := 0; i < len(base); i++ {
		if string(s) == string(base[i]) {
			return i
		}
		index++
	}
	return index
}

func validity_check(str string, base string) bool {
	var checked bool
	if len(base) < 2 {
		return false
	}
	slice_base := []byte(base)
	slice_str := []byte(str)
	for i := 0; i < len(slice_base)-1; i++ {
		if int(slice_base[i]) == 45 || int(slice_base[i]) == 43 {
			return false
		}
		for j := len(slice_base) - 1; j > i; j-- {
			if int(slice_base[j]) == int(slice_base[i]) || int(slice_base[i]) == 45 || int(slice_base[i]) == 43 {
				return false
			}
		}
	}
	for i := 0; i < len(slice_str); i++ {
		checked = false
		for j := 0; j < len(slice_base); j++ {
			if int(slice_str[i]) == int(slice_base[j]) {
				checked = true
				break
			}
		}
		if !checked {
			return false
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	if validity_check(s, base) {
		result := 0
		slice_string := []byte(s)
		reverse_slice_byte_for_atoibase(slice_string)
		for i := 0; i < len(slice_string); i++ {
			result += powforAtoibase(len(base), i) * found_index(slice_string[i], base)
		}
		return result
	}
	return 0
}
