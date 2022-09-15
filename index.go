package piscine

func Index(s string, toFind string) int {
	result := -1
	count := 0
	if len(toFind) == 0 {
		return 0
	}
	if len(toFind) <= len(s) {
		for j := 0; j <= len(s)-len(toFind); j++ {
			count = 0
			if s[j] == toFind[0] {
				for i := 0; i < len(toFind) && s[j+i] == toFind[i]; i++ {
					count = i
				}
				if count == len(toFind)-1 {
					return j
				}
			}
		}
	}
	return result
}
