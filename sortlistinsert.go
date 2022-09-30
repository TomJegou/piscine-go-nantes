package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{Data: data_ref}
	if l.Data >= n.Data {
		n.Next = l
		return n
	}
	current_node := l
	for current_node != nil {
		if current_node.Next.Data >= n.Data {
			n.Next = current_node.Next
			current_node.Next = n
			return l
		}
		current_node = current_node.Next
	}
	return l
}
