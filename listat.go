package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	current_node := l
	for i := 0; i < pos; i++ {
		if current_node.Next == nil {
			return nil
		}
		current_node = current_node.Next
	}
	return current_node
}
