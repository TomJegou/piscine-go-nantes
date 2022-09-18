package piscine

func Split(s string, t string) []string {
	string_temp := ""
	result := []string{}
	counter := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == string(t[counter]) {
			counter++
		} else if string(s[i]) != string(t[counter]) {
			counter = 0
		}
		if counter == len(t) {
			if len(string_temp) > 0 {
				result = append(result, string_temp[:len(string_temp)-1])
				string_temp = ""
				s = s[:i-len(t)] + s[i:]
				i -= len(t)
				counter = 0
			}
		} else {
			string_temp += string(s[i])
		}
	}
	if len(string_temp) > 0 {
		result = append(result, string_temp)
	}
	return result
}
