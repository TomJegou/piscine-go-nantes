package piscine

func ListReverse(l *List) {
	l_temp := &List{}
	current_node := l.Head
	for current_node != nil {
		ListPushFront(l_temp, current_node.Data)
		current_node = current_node.Next
	}
	l.Head = l_temp.Head
	l.Tail = l_temp.Tail
}
