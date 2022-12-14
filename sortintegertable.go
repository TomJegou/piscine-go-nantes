package piscine

func Split_for_sort(t []int) ([]int, int, []int) {
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

func Merge(s_b []int, p int, s_a []int) []int {
	s_b = append(s_b, p)
	s_b = append(s_b, s_a...)
	return s_b
}

func Quick_Sort(t []int) []int {
	if len(t) < 1 {
		return t
	}
	s_1, p, s_2 := Split_for_sort(t)
	a := Merge(Quick_Sort(s_1), p, Quick_Sort(s_2))
	return a
}

func SortIntegerTable(table []int) {
	table_sorted := Quick_Sort(table)
	copy(table, table_sorted)
}
