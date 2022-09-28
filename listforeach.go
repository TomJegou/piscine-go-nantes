package piscine

func ListForEach(l *List, f func(*NodeL)) {
	current_node := l.Head
	for current_node != nil {
		f(current_node)
		current_node = current_node.Next
	}
}
