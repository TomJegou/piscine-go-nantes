package piscine

func PowerForTrimAtoi(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if nb == 1 || nb == 0 {
		return nb
	} else if power == 0 {
		return 1
	}
	a := PowerForTrimAtoi(nb, power/2)
	if power%2 == 0 {
		return a * a
	} else {
		return a * a * nb
	}
}

func isnumForTrimAtoi(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] < '0' || s_S[i] > '9' {
			return false
		}
	}
	return true
}

func convert_to_slice(s string) ([]byte, bool) {
	result := []byte{}
	positiv := true
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "-" {
			positiv = false
		}
		if isnumForTrimAtoi(string(s[i])) {
			result = append(result, byte(s[i])-48)
		}
	}
	return result, positiv
}

func convert_to_int(t []byte, signe bool) int {
	result := 0
	for i := len(t) - 1; i >= 0; i-- {
		result += int(t[i]) * PowerForTrimAtoi(10, len(t)-1-i)
	}
	if signe {
		return result
	} else {
		return result * -1
	}
}

func TrimAtoi(s string) int {
	return convert_to_int(convert_to_slice(s))
}
