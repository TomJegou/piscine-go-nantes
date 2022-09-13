package piscine

func Index(s string, toFind string) int {
	result := -1
	count := 0
	if len(toFind) == 0 {
		return 0
	}
	if len(toFind) <= len(s) {
		for j := 0; j < len(s)-len(toFind); j++ {
			if s[j] == toFind[0] {
				result = j
				count = result
				break
			}
		}
		for i := 0; i < len(toFind); i++ {
			if s[count] == toFind[i] {
				count++
			} else {
				return -1
			}
		}
	}
	return result
}
