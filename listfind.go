package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current_node := l.Head
	for current_node != nil {
		if comp(current_node.Data, ref) {
			return &current_node.Data
		}
		current_node = current_node.Next
	}
	return nil
}
