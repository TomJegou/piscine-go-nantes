package piscine

type NodeL struct {
	Data string
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data string) {
	a := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = a
		l.Tail = a
	}
	l.Tail.Next = a
	l.Tail = l.Tail.Next
}
