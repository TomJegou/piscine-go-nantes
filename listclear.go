package piscine

func ListClear(l *List) {
	l.Head.Next = nil
	l.Head = nil
}
