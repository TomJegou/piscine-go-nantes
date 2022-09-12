package piscine

func StrLen(s string) int {
	var result int
	aS := []byte(s)
	for i := 0; i < len(aS); i++ {
		if aS[i] <= 127 {
			result += 1
		} else {
			continue
		}
	}
	return result
}
