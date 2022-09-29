package piscine

func ListMerge(l1 *List, l2 *List) {
	if ListSize(l1) > 0 {
		l1.Tail.Next = l2.Head
		l2.Head = l1.Head
	} else {
		l1.Head = l2.Head
	}
}
