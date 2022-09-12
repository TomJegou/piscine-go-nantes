package piscine

func Split(t []int) ([]int, int, []int) {
	var index_middle int = len(t) / 2
	pivot := t[index_middle]
	slice := t[:index_middle]
	slice = append(slice, t[index_middle+1:]...)
	slice_before := []int{}
	slice_after := []int{}
	for i := 0; i < len(slice); i++ {
		if slice[i] < pivot {
			slice_before = append(slice_before, slice[i])
		} else {
			slice_after = append(slice_after, slice[i])
		}
	}
	return slice_before, pivot, slice_after
}

func Merge1(s_b []int, p int, s_a []int) []int {
	s_b = append(s_b, p)
	s_b = append(s_b, s_a...)
	return s_b
}

func SortIntegerTable(table []int) []int {
	if len(table) < 1 {
		return table
	}
	s_1, p, s_2 := Split(table)
	return Merge1(SortIntegerTable(s_1), p, SortIntegerTable(s_2))
}
