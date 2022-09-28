package piscine

func ListPushFront(l *List, data interface{}) {
	a := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = a
	} else {
		a.Next = l.Head
		l.Head = a
	}
}
