package piscine

func ListMerge(l1 *List, l2 *List) {
	if ListSize(l1) > 0 && ListSize(l2) > 0 {
		l1.Tail.Next = l2.Head
		l1.Tail = l2.Tail
		l2.Head = l1.Head
	}
}
