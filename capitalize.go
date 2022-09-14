package piscine

func is_lower(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] < 'a' || s_S[i] > 'z' {
			return false
		}
	}
	return true
}

func is_upper(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] < 'A' || s_S[i] > 'Z' {
			return false
		}
	}
	return true
}

func is_num(s string) bool {
	s_S := []byte(s)
	for i := 0; i < len(s_S); i++ {
		if s_S[i] < '0' || s_S[i] > '9' {
			return false
		}
	}
	return true
}

func is_alpha(s string) bool {
	if len(s) == 0 {
		return true
	}
	result := false
	for i := 0; i < len(s); i++ {
		if is_lower(string(s[i])) || is_upper(string(s[i])) || is_num(string(s[i])) {
			result = true
		} else {
			return false
		}
	}
	return result
}

func Capitalize(s string) string {
	slice_string := []byte(s)
	search_mode := true
	for i := 0; i < len(s); i++ {
		if search_mode && is_alpha(string(slice_string[i])) {
			if !is_num(string(slice_string[i])) && !is_upper(string(slice_string[i])) {
				slice_string[i] -= 32
				search_mode = false
			} else {
				search_mode = false
			}
		} else if !is_alpha(string(slice_string[i])) {
			search_mode = true
		} else if !search_mode && is_upper(string(slice_string[i])) {
			slice_string[i] += 32
		}
	}
	return string(slice_string)
}
