package piscine

func Map(f func(int) bool, a []int) []bool {
	result := []bool{}
	for index := 0; index < len(a); index++ {
		result = append(result, f(a[index]))
	}
	return result
}
