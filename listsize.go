package piscine

func ListSize(l *List) int {
	result := 1
	if l.Head == nil {
		return 0
	}
	current_node := l.Head
	for current_node.Next != nil {
		result++
		current_node = current_node.Next
	}
	return result
}
