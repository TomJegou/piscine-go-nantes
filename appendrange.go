package piscine

func AppendRange(min, max int) []int {
	result := []int{}
	if min < max {
		for i := min; i < max; i++ {
			result = append(result, i)
		}
	} else {
		return []int(nil)
	}
	return result
}
